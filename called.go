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

const Doc = "called finds calls specified by the called.funcs flag"

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
		if fn == "" {
			continue
		}

		// split function/method name from the end
		last := strings.LastIndex(fn, ".")
		if last == -1 {
			continue
		}
		prefix := fn[:last]
		name := fn[last+1:]

		// drop surrounding parentheses and detect pointer receiver
		recv := strings.TrimPrefix(prefix, "(")
		recv = strings.TrimSuffix(recv, ")")
		ptr := false
		if strings.HasPrefix(recv, "*") {
			ptr = true
			recv = recv[1:]
		}

		// determine whether it is a method or a package function
		dot := strings.LastIndex(recv, ".")
		slash := strings.LastIndex(recv, "/")
		if dot != -1 && dot > slash {
			// method: pkgpath.Type.Method or (*pkgpath.Type).Method
			pkgpath := recv[:dot]
			typename := recv[dot+1:]
			if ptr {
				typename = "*" + typename
			}
			typ := analysisutil.TypeOf(pass, pkgpath, typename)
			if typ == nil {
				continue
			}
			if m := analysisutil.MethodOf(typ, name); m != nil {
				fs = append(fs, m)
			}
			continue
		}

		// package function: pkgpath.Func
		if ptr {
			// invalid pattern
			continue
		}
		f, _ := analysisutil.ObjectOf(pass, recv, name).(*types.Func)
		if f != nil {
			fs = append(fs, f)
		}
	}

	return fs
}
