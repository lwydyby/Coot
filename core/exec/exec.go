package exec

import (
	"os/exec"
	"bytes"
	"Coot/error"
)

func Execute(shell string) string {
	cmd := exec.Command("/bin/bash", "-c", shell)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	error.Check(err, "执行脚本异常")

	return out.String()
}
