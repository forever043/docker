package main

import (
	"flag"
	"os"
	"testing"

	"github.com/spf13/pflag"
)

var systemTest *bool

func init() {
	cmd := newDaemonCommand()
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		switch f.Value.Type() {
		case "bool":
			flag.Bool(f.Name, false, f.Usage)
			if f.Shorthand != "" {
				flag.Bool(f.Shorthand, false, f.Usage)
			}
			//case "string":
		default:
			flag.String(f.Name, "", f.Usage)
			if f.Shorthand != "" {
				flag.String(f.Shorthand, "", f.Usage)
			}
		}
	})

}

func TestSystem(t *testing.T) {
	if os.Getenv("TEST_ENABLE_COVER") != "" {
		main()
	} else {
		t.Skip("not integration test, check TEST_ENABLE_COVER or TEST_COVERPKG env")
	}
}
