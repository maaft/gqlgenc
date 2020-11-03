package main

import (
	"context"
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/maaft/gqlgenc/clientgen"
	"github.com/maaft/gqlgenc/config"
	"github.com/maaft/gqlgenc/generator"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig(".gqlgenc.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(2)
	}

	clientPlugin := clientgen.New(cfg.Query, cfg.Client)
	if err := generator.Generate(ctx, cfg, api.AddPlugin(clientPlugin)); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(4)
	}
}
