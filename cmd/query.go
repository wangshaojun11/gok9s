/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "查询集群",
	Long:  `查询已经添加到上下文中的k8s集群。`,
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("kubectl", "config", "get-contexts")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr
		e := c.Run()
		if e != nil {
			fmt.Println(e)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
