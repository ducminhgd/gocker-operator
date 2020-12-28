package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	r, err := PruneUnusedImages(cli, ctx)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Space Reclaimded: %d\n", r.SpaceReclaimed)

	for _, dlt := range r.ImagesDeleted {
		fmt.Printf("%s\n", dlt.Deleted)
	}
}

// PruneUnsuedImages prunes images that dangling and built since 4w ago
func PruneUnusedImages(cli *client.Client, ctx context.Context) (types.ImagesPruneReport, error) {
	fargs := filters.NewArgs()
	until := time.Now().AddDate(0, 0, -30).Unix()
	fargs.Add("dangling", "true")
	fargs.Add("until", fmt.Sprintf("%d", until))
	r, err := cli.ImagesPrune(ctx, fargs)
	return r, err
}
