package cmd

import (
	"errors"
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/pkg"
	"os"
	"smart-api/cmd/api"
	"smart-api/cmd/migrate"
	"smart-api/common/global"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "smart",
	Short:        "smart",
	SilenceUsage: true,
	Long:         `smart`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New(pkg.Red("requires at least one arg"))
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 ` + pkg.Green(`smart `+global.Version) + ` 可以使用 ` + pkg.Red(`-h`) + ` 查看命令`
	usageStr1 := `https://blog.csdn.net/weixin_43798031?spm=1000.2115.3001.5343`
	fmt.Printf("%s\n", usageStr)
	fmt.Printf("%s\n", usageStr1)
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(migrate.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
