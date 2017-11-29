// 命令行 
package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var sendCmd = &cobra.Command{
	Use: "send",
	Short: "send block",
	Long: "send block is for new a block and add to blockchain",
}

var queryCmd = &cobra.Command {
	Use: "query",
	Short: "query block",
	Long: "query block from blockchain",
}

var RootCmd = &cobra.Command {
	Use: "bcc",
	Short: "block chain command",
	Long: "block chain command use system",
}

var Flags = &pflag.FlagSet{}

func newStringFlag(name, value, usage string) *pflag.Flag {
	Flags.String(name, value, usage)
	return Flags.Lookup(name)
}

func newIntFlag(name string, value int, usage string) *pflag.Flag {
	Flags.Int(name, value, usage)
	return Flags.Lookup(name)
}

func init() {
	RootCmd.AddCommand(sendCmd)
	RootCmd.AddCommand(queryCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}