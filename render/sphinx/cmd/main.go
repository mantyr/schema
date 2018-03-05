package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/mantyr/schema/json"
	"github.com/mantyr/schema/render/sphinx"
)

func main() {
	src := flag.String("src", "", "source directory")
	dst := flag.String("dst", "", "source directory")
	flag.Parse()

	if *src == "" {
		log.Fatal("empty src")
	}
	if *dst == "" {
		log.Fatal("empty dst")
	}
	err := Render(*src, *dst)
	if err != nil {
		log.Fatal(err)
	}
}
func Render(src, dst string) error {
	files, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, file := range files {
		info, err := os.Stat(
			fmt.Sprintf(
				"%s/%s",
				src, file.Name(),
			),
		)
		if err != nil {
			return err
		}
		if info.IsDir() {
			Render(
				fmt.Sprintf(
					"%s/%s",
					src,
					file.Name(),
				),
				fmt.Sprintf(
					"%s/%s",
					dst,
					file.Name(),
				),
			)
			continue
		}
		srcFile := fmt.Sprintf(
			"%s/%s",
			src,
			file.Name(),
		)
		if filepath.Ext(srcFile) != ".json" {
			continue
		}
		schemas, err := json.NewFileSchemas(srcFile)
		if err != nil {
			return err
		}
		data, err := sphinx.Marshal(schemas)
		if err != nil {
			return err
		}
		ext := path.Ext(file.Name())
		dstFile := fmt.Sprintf(
			"%s/%s.rst",
			dst,
			file.Name()[0:len(file.Name())-len(ext)],
		)
		err = os.MkdirAll(filepath.Dir(dstFile), 0755)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(
			dstFile,
			data,
			0444,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
