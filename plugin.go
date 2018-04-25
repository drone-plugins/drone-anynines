package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (
	Config struct {
		Username     string
		Password     string
		Organization string
		Space        string
		SkipCleanup  bool
	}

	Plugin struct {
		Config Config
	}
)

func (p Plugin) Exec() error {
	cmd := p.command()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	trace(cmd)
	return cmd.Run()
}

func (p *Plugin) command() *exec.Cmd {
	args := []string{
		"--provider=anynines",
	}

	if p.Config.SkipCleanup {
		args = append(
			args,
			"--skip-cleanup",
		)
	}

	args = append(
		args,
		fmt.Sprintf("--username=%s", p.Config.Username),
	)

	args = append(
		args,
		fmt.Sprintf("--password=%s", p.Config.Password),
	)

	args = append(
		args,
		fmt.Sprintf("--organization=%s", p.Config.Organization),
	)

	args = append(
		args,
		fmt.Sprintf("--space=%s", p.Config.Space),
	)

	return exec.Command(
		"dpl",
		args...,
	)
}

func trace(cmd *exec.Cmd) {
	fmt.Println("+", strings.Join(cmd.Args, " "))
}
