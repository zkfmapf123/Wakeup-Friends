package src

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	_defaultProfile = "default"
	cliList         = []string{"dashboard", "wakeup", "sleep", "exit"}
)

var (
	rootCmd = &cobra.Command{
		Use:   "Wake-up-friends",
		Short: "Wake-up-friends",
		Long:  "Wake-up-friends",
		Run: func(cmd *cobra.Command, args []string) {
			Clear()

			profile, region := viper.GetString("profile"), viper.GetString("region")

			cfg, err := config.LoadDefaultConfig(context.TODO(),
				config.WithSharedConfigProfile(profile),
				config.WithRegion(region))

			if err != nil {
				PanicHighLight("aws credentials 지정하고 오셈 default로...")
			}

			for {
				Clear()
				switch selectBox("고르시오", cliList) {
				case "dashboard":
					GetDashboard(cfg)
				case "wakeup":
					BotherEC2(cfg, true)
				case "sleep":
					BotherEC2(cfg, false)
				case "exit":
					fmt.Println("bye")
					os.Exit(0)

				default:
					PanicHighLight("이건 뭐임...")
				}
			}

		},
	}
)

func initial() {

	// init
	cobra.OnInitialize(func() {

		if viper.GetString("profile") == "" {
			viper.Set("profile", "default")
		}

		if viper.GetString("region") == "" {
			viper.Set("region", "ap-northeast-2")
		}
	})

	rootCmd.PersistentFlags().StringP("profile", "p", "", "[Optional] aws profile")
	rootCmd.PersistentFlags().StringP("region", "r", "", "[Optional] aws region")

	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("region", rootCmd.PersistentFlags().Lookup("region"))
}

func Execute() {

	Clear()
	initial()
	if err := rootCmd.Execute(); err != nil {
		PanicHighLight(err.Error())
	}
}

func PanicHighLight(err string) {
	fmt.Println(color.RedString("에러요 : ", errors.New(err)))
	os.Exit(1)
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
