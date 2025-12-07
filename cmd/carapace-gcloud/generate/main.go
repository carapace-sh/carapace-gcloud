package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	spec "github.com/carapace-sh/carapace-spec"
	"gopkg.in/yaml.v3"
)

func main() {
	services, err := services("../cmd/gcloud")
	if err != nil {
		panic(err.Error())
	}
	err = os.WriteFile("../cmd/gcloud/gcloud_generated.go", []byte(services), os.ModePerm)
	if err != nil {
		panic(err.Error())
	}
	err = execCommand("go", "fmt", "../cmd/gcloud/gcloud_generated.go")
	if err != nil {
		panic(err.Error())
	}
}

func execCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	println("# " + strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func services(dir string) (string, error) {
	r := regexp.MustCompile(`^gcloud\.(?P<service>[^.]+)\.yaml$`)
	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}

	s := `package gcloud

func init() {
	services = map[string]string{
`

	for _, entry := range entries {
		if !r.MatchString(entry.Name()) {
			continue
		}

		content, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			return "", err
		}

		var command spec.Command
		if err := yaml.Unmarshal(content, &command); err != nil {
			return "", err
		}

		s += fmt.Sprintf("\t\t%q: %q,\n", command.Name, command.Description)
	}

	s += `	}
}
`
	return s, nil
}
