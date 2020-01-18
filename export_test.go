package called

func ExportSetFlagFuncs(s string) func() {
	org := flagFuncs
	flagFuncs = s
	return func() {
		flagFuncs = org
	}
}
