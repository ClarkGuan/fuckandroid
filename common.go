package fa

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
)

var errNoWorkspace = errors.New("no workspace directory found")
var errAlreadyExist = errors.New("path already exist")

func makeDirs(dirs ...string) error {
	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0775); err != nil {
			return err
		}
	}
	return nil
}

func boxCopyAll(list []*struct {
	from string
	to   string
	perm os.FileMode
}, box *rice.Box) error {
	for _, entry := range list {
		if err := boxCopy(box, entry.from, entry.to, entry.perm); err != nil {
			return err
		}
	}
	return nil
}

func boxCopy(box *rice.Box, from, to string, perm os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(to), 0775); err != nil {
		return err
	}
	return ioutil.WriteFile(to, box.MustBytes(from), perm)
}

func boxCopyTemplate(box *rice.Box, from, to string, perm os.FileMode, holder map[string]string) error {
	if err := os.MkdirAll(filepath.Dir(to), 0775); err != nil {
		return err
	}

	tmpl := template.New(filepath.Base(from))
	tmpl, _ = tmpl.Parse(box.MustString(from))
	output, err := os.OpenFile(to, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer output.Close()
	return tmpl.Execute(output, holder)
}
