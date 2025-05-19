package called

import (
	"testing"

	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestRestrictedFuncsDuplicate(t *testing.T) {
	testdata := analysistest.TestData()
	var funcs []*types.Func
	a := &analysis.Analyzer{
		Name: "test",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			funcs = restrictedFuncs(pass, "b.Func,b.Func,(*b.Type).Method,(*b.Type).Method,b/bsub.Type.Method,b/bsub.Type.Method")
			return nil, nil
		},
	}
	analysistest.Run(t, testdata, a, "a")
	if len(funcs) != 3 {
		t.Fatalf("expected 3 funcs, got %d", len(funcs))
	}
}
