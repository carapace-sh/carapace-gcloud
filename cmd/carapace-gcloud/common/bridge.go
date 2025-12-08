package common

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
)

func ActionBridgeGcloudCompleter() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		c.Args = carapace.NewContext(os.Args[4:]...).Args // TODO nasty args passthrough
		return bridge.ActionGcloud("gcloud").Invoke(c).ToA()
	})
}
