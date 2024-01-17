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

			cfg, err := config.LoadDefaultConfig(context.TODO())
			if err != nil {
				PanicHighLight("aws credentials 지정하고 오셈 default로...")
			}

			for {
				Clear()
				switch selectBox("고르시오", cliList) {
				case "dashboard":
					GetDashboard(cfg)
				case "wakeup":
					ExecuteWakeup(cfg)
				case "sleep":
					ExecuteSleep(cfg)
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

func Execute() {

	Clear()
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
