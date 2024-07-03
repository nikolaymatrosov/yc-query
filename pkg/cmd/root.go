package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/expr-lang/expr"
	"github.com/mitchellh/go-homedir"
	"github.com/nikolaymatrosov/yc-query/pkg/config"
	"github.com/nikolaymatrosov/yc-query/pkg/query"
	"github.com/spf13/cobra"
	ycsdk "github.com/yandex-cloud/go-sdk"
)

// RootCommand returns the root command for the CLI
type RootCommand struct {
	configPath string
	input      string
}

func Root() *cobra.Command {
	var rootCmd = &RootCommand{}
	var cmd = &cobra.Command{
		Use:   "yc-query",
		Short: "Fluent Yandex.Cloud SDK wrapper",
		Long:  `yc-query is a wrapper around the Yandex.Cloud SDK that provides a fluent API for querying and manipulating resources.`,
		Run:   rootCmd.Run,
	}

	cmd.PersistentFlags().StringVarP(&rootCmd.configPath, "config", "c", "~/.config/yandex-cloud/config.yaml", "config file (default is ~/.config/yandex-cloud/config.yaml)")
	cmd.PersistentFlags().StringVarP(&rootCmd.input, "input", "i", "", "input file (default is stdin)")

	return cmd
}

func (r *RootCommand) Run(cmd *cobra.Command, args []string) {
	ctx := cmd.Context()
	configPath, err := homedir.Expand(r.configPath)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	conf, err := config.Parse(configPath)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}
	profile := conf.CurrentProfile()

	var cred ycsdk.Credentials
	if profile.Token != "" {
		cred = ycsdk.OAuthToken(profile.Token)
	} else if profile.ServiceAccountKey.Id != "" {
		cred, err = ycsdk.ServiceAccountKey(profile.ServiceAccountKey.IamKey())
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	} else {

	}

	sdk, err := ycsdk.Build(ctx, ycsdk.Config{
		Credentials: cred,
	})
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	var input *os.File

	if r.input != "" && r.input != "-" {
		input, err = os.Open(r.input)
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
	} else {
		input = os.Stdin
	}

	code, err := io.ReadAll(input)
	if err != nil {
		cmd.PrintErrln(err)
		return
	}

	env := query.NewEnv(ctx, sdk)

	program, err := expr.Compile(string(code), expr.Env(env))
	if err != nil {
		panic(err)
	}

	output, err := expr.Run(program, env)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

}
