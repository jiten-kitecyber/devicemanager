package cmd

import (
	"github.com/jiten-kitecyber/devicemanager/manager"

	"log"

	"github.com/spf13/cobra"
)

var configType string
var scope string
var primaryDNS string
var secondaryDNS string
var iface string
var show string

func init() {

	configureCmd.Flags().StringVarP(&configType, "type", "t", "dns", "config type.dns|firewall")
	configureCmd.Flags().StringVarP(&scope, "scope", "s", "system", "two types of the scopes. system|command")
	configureCmd.Flags().StringVarP(&primaryDNS, "primarydns", "", "", "provide primary dns")
	configureCmd.Flags().StringVarP(&secondaryDNS, "secondarydns", "", "", "provide secondary dns")
	configureCmd.Flags().StringVarP(&iface, "interface", "i", "all", "provide interfaces based on the system")
	configureCmd.Flags().StringVarP(&show, "show", "", "", "shows current dns|firewall settings")

	rootCmd.AddCommand(configureCmd)
}

var configureCmd = &cobra.Command{
	Use:   "config",
	Short: "config is to configure dns and firewall",
	Long:  `config is to configure dns and firewall.dns or firewall settings to be supplied`,
	Run: func(cmd *cobra.Command, args []string) {
		if configType == "dns" { // system wide dns
			var idm manager.IDNSDeviceManager
			if scope == "system" {
				idm = new(manager.GlobalDNS)
				if primaryDNS == "" || secondaryDNS == "" {
					log.Fatalln("primary and secondary dns ips must be given")
				}
				err := idm.SetDNS("", primaryDNS, secondaryDNS)
				if err != nil {
					log.Fatalln(err)
				}
			} else if scope == "command" {
				if iface == "" {
					log.Fatalln("interface cannot be empty")
				}
				idm = new(manager.CommandDNS)
				if primaryDNS == "" || secondaryDNS == "" {
					log.Fatalln("primary and secondary dns ips must be given")
				}
				err := idm.SetDNS(iface, primaryDNS, secondaryDNS)
				if err != nil {
					log.Fatalln(err)
				}

			}
		}
	},
}
