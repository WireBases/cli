package cmd

import (
	"fmt"
	Pkg "wirebasec/Pkg"

	"github.com/spf13/cobra"
)

var mysql = &cobra.Command{
	Use:   "mysql",
	Short: "You can check mysql",
	Long:  `Where you can control with operations such as opening and closing Mysql`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("ERROR: Undefiend arguments")
			fmt.Println("Use command wirebasec mysql start/stop")
			return
		}

		argsr := args[0]

		if argsr == "start" {
			_, strng := Pkg.Mysql_start()
			fmt.Println(strng)
		}

		if argsr == "stop" {
			_, strng := Pkg.Mysql_stop()
			fmt.Println(strng)
		}

		if argsr != "" && argsr != "start" && argsr != "stop" {
			fmt.Println("ERROR:", "Invalid argument")
			fmt.Println("Arguments must be one of use wirebasec mysql start, stop")
		}
	},
}

func init() {
	rootCmd.AddCommand(mysql)
}
