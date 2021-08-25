package command

import (
	"errors"
	"fmt"
	"strings"
)

func updateEnvContext(args []string, ctx *Context) (err error) {
	for i := 0; i < len(args); i++ {
		envar := strings.SplitN(args[i], "=", 2)
		if len(envar) == 2 {
			key := envar[0]
			value := envar[1]
			ctx.Env[key] = value
		} else {
			err = errors.New(fmt.Sprintf("env: ‘%s’: No such file or directory", args[i]))
		}
	}
	return nil
}
