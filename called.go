package called

import (
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

var flagFuncs string

var Analyzer = &analysis.Analyzer{
	Name: "called",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func init() {
	Analyzer.Flags.StringVar(&flagFuncs, "funcs", "", "function or method names which are restricted calling")
}

const Doc = "called find callings specified by called.funcs flag"

func run(pass *analysis.Pass) (interface{}, error) {
	if flagFuncs == "" {
		return nil, nil
	}

	fs := restrictedFuncs(pass, flagFuncs)
	if len(fs) == 0 {
		return nil, nil
	}

	pass.Report = analysisutil.ReportWithoutIgnore(pass)
	srcFuncs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs
	for _, sf := range srcFuncs {
		for _, b := range sf.Blocks {
			for _, instr := range b.Instrs {
				for _, f := range fs {
					if analysisutil.Called(instr, nil, f) {
						pass.Reportf(instr.Pos(), "%s must not be called", f.FullName())
						break
					}
				}
			}
		}
	}

	return nil, nil
}

func restrictedFuncs(pass *analysis.Pass, names string) []*types.Func {
	var fs []*types.Func
	for _, fn := range strings.Split(names, ",") {
		fn = strings.TrimSpace(fn)
		if len(fn) == 0 {
			continue
		}

		if fn[0] == '(' {
			// method: (*pkgname.Type).Method
			ss := splitLastN(fn, ".", 3)
			if len(ss) < 3 {
				continue
			}
			pkgname := strings.TrimLeft(ss[0], "(")
			typename := strings.TrimRight(ss[1], ")")
			if pkgname != "" && pkgname[0] == '*' {
				pkgname = pkgname[1:]
				typename = "*" + typename
			}

			typ := analysisutil.TypeOf(pass, pkgname, typename)
			if typ == nil {
				continue
			}

			m := analysisutil.MethodOf(typ, ss[2])
			if m != nil {
				fs = append(fs, m)
			}
		} else {
			// package function: pkgname.Func
			ss := splitLastN(fn, ".", 2)
			if len(ss) < 2 {
				continue
			}
			f, _ := analysisutil.ObjectOf(pass, ss[0], ss[1]).(*types.Func)
			if f != nil {
				fs = append(fs, f)
				continue
			}
		}
	}

	return fs
}

func splitLastN(s, sep string, n int) []string {
	ret := make([]string, 0)
	for i := 0; i < n-1; i++ {
		li := strings.LastIndex(s, sep)
		if li < 0 {
			break
		}

		if li+1 < len(s) {
			ret = append(ret, s[li+1:])
		} else {
			ret = append(ret, "")
		}
		s = s[:li]
	}
	ret = append(ret, s)

	// reverse
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}
