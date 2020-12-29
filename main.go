package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	// Prune command
	pruneCmd := flag.NewFlagSet("prune", flag.ExitOnError)
	danglingPtr := pruneCmd.Bool("dangling", true, "Remove dangling images. Default: true")
	untilPtr := pruneCmd.Int("until", 30, "Keep images are updated from `until` day(s) before.")
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Command is require: prune")
		os.Exit(1)
	}

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	switch (flag.Arg(0)) {
	case "prune":
		r, err := PruneUnusedImages(cli, ctx, *danglingPtr, *untilPtr)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\tSpace Reclaimded: %d\n", r.SpaceReclaimed)

		if len(r.ImagesDeleted) > 0 {
			fmt.Println("\tDeleted")
		}
		for _, dlt := range r.ImagesDeleted {
			fmt.Printf("\t\t%s\n", dlt.Deleted)
		}
		break
	default:
		allowedCmds := []string {
			"prune",
		}
		fmt.Printf("Allowed commands is: %s\n", strings.Join(allowedCmds, ", "))
		os.Exit(1)
	}
	os.Exit(0)
}

// PruneUnsuedImages prunes images that dangling and built since 4w ago
func PruneUnusedImages(cli *client.Client, ctx context.Context, dlg bool, days int) (types.ImagesPruneReport, error) {
	fargs := filters.NewArgs()
	if dlg {
		fargs.Add("dangling", "true")
	} else {
		fargs.Add("dangling", "false")
	}
	until := time.Now().AddDate(0, 0, -days).Unix()
	fargs.Add("until", fmt.Sprintf("%d", until))
	r, err := cli.ImagesPrune(ctx, fargs)
	return r, err
}
