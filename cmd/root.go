/*
Copyright © 2023 CFET CFET CFET

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/csid-cfet/shuffler/cmd/shuffle"
	"github.com/csid-cfet/shuffler/cmd/team"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/common-nighthawk/go-figure"
)

var cfgFile string


// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shuffler",
	Short: "Shuffle the team.",
	Long: "\n" + figure.NewFigure("Shuffler", "nancyj-underlined", true).String(),
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func addSubCommandPalettes() {
	rootCmd.AddCommand(shuffle.ShuffleCmd)
	rootCmd.AddCommand(team.TeamCmd)
}


func init() {
	cobra.OnInitialize(initConfig)

	addSubCommandPalettes()
}


// TODO: Find an actual use for Viper.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".shuffler" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".shuffler")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
