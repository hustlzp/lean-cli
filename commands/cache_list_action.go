package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/leancloud/lean-cli/api"
	"github.com/leancloud/lean-cli/apps"
	"github.com/urfave/cli"
)

func cacheListAction(c *cli.Context) error {
	appID, err := apps.GetCurrentAppID(".")
	if err != nil {
		return err
	}

	caches, err := api.GetCacheList(appID)
	if err != nil {
		return err
	}

	if len(caches) == 0 {
		return cli.NewExitError("This app doesn't have any LeanCache instance", 1)
	}

	t := tabwriter.NewWriter(os.Stdout, 0, 1, 3, ' ', 0)

	fmt.Fprintln(t, "InstanceName\tMaxMemory")
	for _, cache := range caches {
		fmt.Fprintf(t, "%s\t%dM\r\n", cache.Instance, cache.MaxMemory)
	}
	t.Flush()

	return nil
}
