//#go:build integration

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-gcloud/cmd/carapace-gcloud/cmd/gcloud"
)

func reportPatch(t *testing.T, title string, patch []string) {
	t.Helper()

	s := []string{fmt.Sprintf("\033[2m# %v\033[0m", title)}
	for _, line := range patch {
		switch {
		case strings.HasPrefix(line, "-"):
			s = append(s, fmt.Sprintf("\033[0;31m%v\033[0m", line))
			t.Fail()
		case strings.HasPrefix(line, "+"):
			s = append(s, fmt.Sprintf("\033[0;32m%v\033[0m", line))
			t.Fail()
		}
	}
	fmt.Println(strings.Join(s, "\n"))
}

func TestService(t *testing.T) {
	testDir, err := os.MkdirTemp("", "carapace-gcloud_testService-*")
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(testDir)
	testFile, err := os.Create(filepath.Join(testDir, "outfile"))
	if err != nil {
		t.Fatal(err.Error())
	}
	defer os.Remove(testFile.Name())

	serviceToTest := os.Getenv("SERVICE")

	if serviceToTest == "" || serviceToTest == "_" {
		fmt.Printf("\033[2m# %v\033[0m\n", "_")

		patch := carapace.DiffPatch(
			bridge.ActionGcloud("gcloud"),
			bridge.ActionCarapace("carapace-gcloud").Chdir(testDir),
			carapace.NewContext(""),
		)
		reportPatch(t, "_", patch)
	}

	for service := range gcloud.Services() {
		t.Run(service, func(t *testing.T) {
			if serviceToTest != "" && service != serviceToTest {
				t.SkipNow()
			}

			fmt.Printf("\033[2m# %v\033[0m\n", service)

			patch := carapace.DiffPatch(
				bridge.ActionGcloud("gcloud"),
				bridge.ActionCarapace("carapace-gcloud").Chdir(testDir),
				carapace.NewContext(service, ""),
			)
			reportPatch(t, service, patch)

			command, err := gcloud.Get(fmt.Sprintf("gcloud.%s.yaml", service))
			if err != nil {
				t.Fatal(err.Error())
			}
			for _, operation := range command.Commands {
				t.Run(operation.Name, func(t *testing.T) {
					t.Parallel()
					patch := carapace.DiffPatch(
						bridge.ActionGcloud("gcloud"),
						bridge.ActionCarapace("carapace-gcloud").Chdir(testDir),
						carapace.NewContext(service, operation.Name, "--"),
					)

					reportPatch(t, fmt.Sprintf("%v %v", service, operation.Name), patch)

					if len(operation.Commands) == 0 {
						return
					}

					patch = carapace.DiffPatch(
						bridge.ActionGcloud("gcloud"),
						bridge.ActionCarapace("carapace-gcloud").Chdir(testDir),
						carapace.NewContext(service, operation.Name, ""),
					)
					reportPatch(t, fmt.Sprintf("%v %v", service, operation.Name), patch)

					for _, subOperation := range operation.Commands {
						t.Run(subOperation.Name, func(t *testing.T) {
							t.Parallel()
							patch := carapace.DiffPatch(
								bridge.ActionGcloud("gcloud"),
								bridge.ActionCarapace("carapace-gcloud").Chdir(testDir),
								carapace.NewContext(service, operation.Name, subOperation.Name, "--"),
							)

							reportPatch(t, fmt.Sprintf("%v %v %v", service, operation.Name, subOperation.Name), patch)

							if len(subOperation.Commands) == 0 {
								return
							}

							patch = carapace.DiffPatch(
								bridge.ActionGcloud("gcloud"),
								bridge.ActionCarapace("carapace-gcloud").Chdir(testDir),
								carapace.NewContext(service, operation.Name, subOperation.Name, ""),
							)
							reportPatch(t, fmt.Sprintf("%v %v %v", service, operation.Name, subOperation.Name), patch)

							for _, subSubOperation := range subOperation.Commands {
								t.Run(subSubOperation.Name, func(t *testing.T) {
									t.Parallel()
									patch := carapace.DiffPatch(
										bridge.ActionGcloud("gcloud"),
										bridge.ActionCarapace("carapace-gcloud").Chdir(testDir),
										carapace.NewContext(service, operation.Name, subOperation.Name, subSubOperation.Name, "--"),
									)

									reportPatch(t, fmt.Sprintf("%v %v %v %v", service, operation.Name, subOperation.Name, subSubOperation.Name), patch)
								})
							}
						})
					}
				})
			}
		})
	}
}
