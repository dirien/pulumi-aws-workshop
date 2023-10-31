package main

import (
	"github.com/pulumi/pulumi-docker/sdk/v4/go/docker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		cfg := config.New(ctx, "")

		image, err := docker.NewImage(ctx, "image", &docker.ImageArgs{
			Build: docker.DockerBuildArgs{
				Context:  pulumi.String("./app"),
				Platform: pulumi.String("linux/amd64"),
			},
			ImageName: pulumi.Sprintf("%s/myapp:latest", cfg.Get("registry")),
			Registry: docker.RegistryArgs{
				Server:   pulumi.String(cfg.Get("registry")),
				Username: pulumi.String(cfg.Get("username")),
				Password: cfg.GetSecret("password"),
			},
		})
		if err != nil {
			return err
		}
		ctx.Export("imageDigest", image.RepoDigest)

		return nil
	})
}
