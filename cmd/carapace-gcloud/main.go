package main

import "github.com/carapace-sh/carapace-gcloud/cmd/carapace-gcloud/cmd"

//go:generate sh -c "go run -C ./generate ."
func main() {
	cmd.Execute()
}
