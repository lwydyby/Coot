package file

import (
	"fmt"
	"os"
	"Coot/error"
)

func Output(result string, path string) {
	if (path != "") {
		_, err := os.Stat(path)
		if err != nil {
			f_create, err_create := os.Create(path)
			error.Check(err_create, "File creation failure")
			f_create.Close()
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
		error.Check(err, "fail to open file")
		f.Write([]byte(result))
		f.Close()
	} else {
		fmt.Println(result)
	}
}
