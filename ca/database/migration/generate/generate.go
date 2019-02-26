package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	fileHeader := "package migration\n\nfunc Files() *map[string]string {\n\n\tfiles := make(map[string]string)\n\n"
	fileFooter := "\n\treturn &files\n}\n"

	dir, _ := os.Getwd()
	sqlPath := filepath.Join(dir, "ca", "database", "migration", "sql")
	outPath := filepath.Join(dir, "ca", "database", "migration")

	fs, _ := ioutil.ReadDir(sqlPath)

	out, _ := os.Create(filepath.Join(outPath, "migration.go"))
	out.Write([]byte(fileHeader))

	for _, f := range fs {

		if strings.HasSuffix(f.Name(), ".sql") {

			data := "\tfiles[\""
			data += f.Name()
			data += "\"] = \""

			fileData, err := ioutil.ReadFile(filepath.Join(sqlPath, f.Name()))
			if err != nil {
				fmt.Print(err.Error())
				return
			}

			a := make([]byte, 1)
			for _, v := range fileData {
				a[0] = v
				data += "\\x"
				data += hex.EncodeToString(a)
			}

			out.Write([]byte(data))
			out.Write([]byte("\"\n"))
		}
	}

	out.Write([]byte(fileFooter))
}
