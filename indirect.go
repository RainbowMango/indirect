package indirect

import (
	"github.com/Rican7/retry"
	"github.com/spf13/pflag"
)

var ip *int = pflag.Int("flagname", 1234, "help message for flagname")

func GetFlagValue() int {
	return *ip
}

func RetryFunc() {
	retry.Retry(func(attempt uint) error {
		return nil // Do something that may or may not cause an error
	})
}
