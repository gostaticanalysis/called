package called_test

import (
	"testing"

	"github.com/gostaticanalysis/called"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	defer called.ExportSetFlagFuncs("b.Func,(*b.Type).Method, (b/bsub.Type).Method,github.com/gostaticanalysis/c.Func,(*github.com/gostaticanalysis/c.Type).Method,(github.com/gostaticanalysis/c/csub.Type).Method")()
	analysistest.Run(t, testdata, called.Analyzer, "a")
}
