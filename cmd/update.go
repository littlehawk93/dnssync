/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/littlehawk93/dnssync/icanhazip"
	"github.com/spf13/cobra"
)

var updateDomains []string

var updateProvider string

var updateForce bool

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates one or more DNS entries with the latest IP",
	Run:   runUpdate,
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringSliceVarP(&updateDomains, "domains", "d", nil, "domains to update")
	updateCmd.Flags().StringVarP(&updateProvider, "provider", "p", "", "DNS provider to use to update")
	updateCmd.Flags().BoolVarP(&updateForce, "force", "f", false, "update IP address even if IP hasn't changed")

	updateCmd.MarkFlagRequired("domains")
	updateCmd.MarkFlagRequired("provider")
}

func runUpdate(cmd *cobra.Command, args []string) {

	prov := configuration.GetMatchingProvider(updateProvider)

	if prov == nil {
		log.Fatalf("invalid provider name '%s'", updateProvider)
	}

	ip, err := icanhazip.GetIP()

	if err != nil {
		log.Fatalf("failed to get IP address: %s", err.Error())
	}

	if ip == nil {
		log.Fatal("invalid IP address format")
	}

	for _, domain := range updateDomains {
		if err := prov.UpdateIP(ip, domain, 3600, updateForce); err != nil {
			log.Fatalf("failed to update domain '%s': %s", domain, err.Error())
		}
	}
}
