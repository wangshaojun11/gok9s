/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// switchCmd represents the switch command
var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "切换集群",
	Long:  `使用 switch 子命令, 可以切换集群, 并且进入k9s。`,
	Run: func(cmd *cobra.Command, args []string) {
		c, _ := cmd.Flags().GetString("cluster")
		n, _ := cmd.Flags().GetString("namespace")
		c2 := exec.Command(
			"k9s",
			"--context",
			c,
			"-n",
			n)
		e := c2.Run()
		if e != nil {
			fmt.Println(e)
		}
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// switchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	switchCmd.Flags().StringP("cluster", "c", "ucloud-test", "指定切换的集群")
	switchCmd.Flags().StringP("namespace", "n", "default", "指定切换的命名空间")

}
