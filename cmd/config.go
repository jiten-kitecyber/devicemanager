package cmd

import (
	"github.com/spf13/cobra"
)

var configType string
var scope string
var primaryDNS string
var secondaryDNS string
var iface string

func init() {

	configureCmd.Flags().StringVarP(&configType, "type", "t", "dns", "config type.dns|firewall")
	configureCmd.Flags().StringVarP(&scope, "scope", "s", "system", "two types of the scopes. system|command")
	configureCmd.Flags().StringVarP(&primaryDNS, "primarydns", "pd", "", "provide primary dns")
	configureCmd.Flags().StringVarP(&secondaryDNS, "secondarydns", "pd", "", "provide secondary dns")
	configureCmd.Flags().StringVarP(&iface, "interface", "i", "", "provide interfaces based on the system")

	rootCmd.AddCommand(configureCmd)
}

var configureCmd = &cobra.Command{
	Use:   "config",
	Short: "config is to configure dns and firewall",
	Long:  `config is to configure dns and firewall.dns or firewall settings to be supplied`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}
