package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xfhelp",
	Short: "A CLI tool for helping with XFCE configuration",
}

type queryFunc func(args ...string) ([]byte, error)

func realFunc(args ...string) ([]byte, error) {
	cmd := exec.Command("xfconf-query", args...)
	cmd.Env = append(os.Environ(), "LC_ALL=C")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("xfconf-query: %w", err)
	}
	return output, nil
}

func Execute(version string) error {
	rootCmd.Version = version
	return rootCmd.Execute()
}
