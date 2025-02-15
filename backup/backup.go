package backup

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexpfx/go-pass_manager/pass"
)

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return homeDir
}

func Export(path string) error {
	passDir := getHomeDir() + "/.password-store"
	outputFile := path

	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	passlist, err := pass.List(passDir)
	if err != nil {
		return fmt.Errorf("failed to list passwords: %w", err)
	}
	for _, passName := range passlist {
		password, err := pass.Show(passName)
		if err != nil {
			continue
		}
		s := fmt.Sprintf("%s: %s", passName, password)
		_, err = file.WriteString(s + "\n")
		if err != nil {
			return fmt.Errorf("failed to write to output file: %w",
				err)
		}

	}

	fmt.Printf("Passwords exported to %s\n", outputFile)
	return nil
}

func Load(path string) {
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
