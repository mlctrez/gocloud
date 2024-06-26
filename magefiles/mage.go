package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

var Default = Build

func Build(ctx context.Context) error {
	return run(ctx, makeTemp, buildWasm)
}

func makeTemp(ctx context.Context) (err error) {
	return os.MkdirAll("temp", 0755)
}

func buildWasm(ctx context.Context) error {
	return goCmd(false, "build", "-o", "temp/gocloud", "main.go")
}

func goCmd(wasm bool, args ...string) error {
	cmd := exec.Command("go", args...)
	cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
	output, err := cmd.CombinedOutput()
	if len(output) > 0 {
		fmt.Println(string(output))
	}
	return err
}

func run(ctx context.Context, commands ...func(ctx context.Context) error) error {
	for _, command := range commands {
		if err := command(ctx); err != nil {
			return err
		}
	}
	return nil
}
