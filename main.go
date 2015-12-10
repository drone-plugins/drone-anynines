package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Anynines Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Username) == 0 {
		fmt.Println("Please provide a username")

		os.Exit(1)
		return
	}

	if len(vargs.Password) == 0 {
		fmt.Println("Please provide a password")

		os.Exit(1)
		return
	}

	if len(vargs.Organization) == 0 {
		fmt.Println("Please provide an organization")

		os.Exit(1)
		return
	}

	if len(vargs.Space) == 0 {
		fmt.Println("Please provide a space")

		os.Exit(1)
		return
	}

	dpl := buildDpl(&workspace, &repo, &build, &vargs)

	dpl.Dir = workspace.Path
	dpl.Stderr = os.Stderr
	dpl.Stdout = os.Stdout

	trace(dpl)

	if err := dpl.Run(); err != nil {
    fmt.Println(err)

		os.Exit(1)
		return
	}
}

func buildDpl(workspace *drone.Workspace, repo *drone.Repo, build *drone.Build, vargs *Params) *exec.Cmd {
	args := []string{
		"--provider=anynines",
	}

	args = append(args, fmt.Sprintf(
		"--username=%s",
		vargs.Username))

	args = append(args, fmt.Sprintf(
		"--password=%s",
		vargs.Password))

	args = append(args, fmt.Sprintf(
		"--organization=%s",
		vargs.Organization))

	args = append(args, fmt.Sprintf(
		"--space=%s",
		vargs.Space))

	return exec.Command("dpl", args...)
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
