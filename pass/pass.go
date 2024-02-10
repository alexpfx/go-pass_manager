package pass

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
)

func Insert(password, name string, force bool) error {
	_, err := Show(name)
	if err == nil && !force {
		return fmt.Errorf("a senha com o nome %s já existe. para substituir use a flag --force", name)
	}

	return insert(password, name)
}

func Show(name string) (string, error) {
	p := script.Exec(fmt.Sprintf("pass %s", name))
	str, err := p.String()
	return strings.TrimSpace(str), err
}

func List(storeDir string) ([]string, error) {
	list := readPassList(storeDir)
	return list, nil
}

func readPassList(storeDir string) []string {
	paths := make([]string, 0)
	filepath.WalkDir(storeDir, func(path string, dirEntry fs.DirEntry, err error) error {
		if dirEntry.IsDir() {
			return nil
		}

		ext := filepath.Ext(path)
		if ext != ".gpg" {
			return nil
		}

		name := strings.Replace(path, fmt.Sprintf("%s%s", storeDir, string(filepath.Separator)), "", 1)

		name = strings.Replace(name, ext, "", 1)
		name = strings.ReplaceAll(name, string(filepath.Separator), "/")

		paths = append(paths, name)
		return nil
	})

	return paths

}

func insert(password, name string) error {
	_, err := script.Echo(password).Exec(fmt.Sprintf("pass insert --echo %s", name)).String()
	return err
}
