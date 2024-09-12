/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/littlehawk93/dnssync/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

var configuration config.Configuration

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dnssync",
	Short: "Tools for syncing DNS records",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Configuration file")

	rootCmd.MarkPersistentFlagFilename("config")
	rootCmd.MarkPersistentFlagRequired("config")

	cobra.OnInitialize(loadConfigurationFile)
}

func loadConfigurationFile() {
	if strings.TrimSpace(configFile) == "" {
		log.Fatal("no config file provided")
	}

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(fmt.Errorf("error reading config file '%s': %s", configFile, err.Error()))
	}

	configuration = config.Configuration{}

	if err := viper.Unmarshal(&configuration); err != nil {
		log.Fatal(fmt.Errorf("error parsing config file '%s': %s", configFile, err.Error()))
	}
}
