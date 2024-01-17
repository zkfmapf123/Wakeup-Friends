package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

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

			switch selectBox("고르시오", cliList) {

			case "dashboard":
				GetDashboard()
			case "wakeup":
				ExecuteWakeup()
			case "sleep":
				ExecuteSleep()
			case "exit":
				fmt.Println("bye")
				os.Exit(0)

			default:
				PanicHighLight(errors.New("이건 뭐임..."))
			}
		},
	}
)

func Execute() {

	Clear()
	if err := rootCmd.Execute(); err != nil {
		PanicHighLight(err)
	}
}

func PanicHighLight(err error) {
	fmt.Println(color.RedString("에러요 : ", err.Error()))
	os.Exit(1)
}

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
