// 命令行 
package main

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "add block",
	Long: "add block is for new a block and add to blockchain",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var printCmd = &cobra.Command{
	Use: "print",
	Short: "print all block in blockchain",
	Long: "Print all block in blockchain with current dot",
}

func init() {
	
}

func Execute() {
	
}