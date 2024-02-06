package dm

import (
	"errors"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type CommandDNS struct{}

func (cd *CommandDNS) SetDNS(primaryDNS, secondaryDNS string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		if !cd.HasCommand("netsh") {
			return fmt.Errorf("netsh command not found for operating system: %s", runtime.GOOS)
		}
		interfaceNames, err := cd.GetActiveInterfaces()
		if interfaceNames == nil {
			return fmt.Errorf("unable to determine active interface")
		} else if err != nil {
			return err
		}
		for _, ifacename := range interfaceNames {
			cmd := exec.Command("netsh", "interface", "ip", "set", "dns", "name="+ifacename, "static", primaryDNS, secondaryDNS)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

	case "linux":
		if !cd.HasCommand("nmcli") {
			return fmt.Errorf("nmcli command not found, consider installing NetworkManager or use an alternative method for your Linux distribution")
		}

		interfaceNames, err := cd.GetActiveInterfaces()
		if interfaceNames == nil {
			return fmt.Errorf("unable to determine active interface")
		} else if err != nil {
			return err
		}

		for _, ifacename := range interfaceNames {
			cmd = exec.Command("nmcli", "connection", "modify", ifacename, "ipv4.dns", strings.Join([]string{primaryDNS, secondaryDNS}, ","))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}

	case "darwin":
		if !cd.HasCommand("networksetup") {
			return fmt.Errorf("networksetup command not found, consider installing it")
		}
		interfaceNames, err := cd.GetActiveInterfaces()
		if interfaceNames == nil {
			return fmt.Errorf("unable to determine active interface")
		} else if err != nil {
			return err
		}
		for _, ifacename := range interfaceNames {
			cmd := exec.Command("networksetup", "-setdnsservers", ifacename, primaryDNS, secondaryDNS)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
		}
	default:
		return fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	return cmd.Run()
}

func (cd *CommandDNS) GetActiveInterfaces() ([]string, error) {
	ifaces := make([]string, 0)
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, errors.New("Error getting network interfaces:" + err.Error())

	}

	for _, iface := range interfaces {
		if (iface.Flags&net.FlagUp) != 0 && (iface.Flags&net.FlagLoopback) == 0 {
			ifaces = append(ifaces, iface.Name)
		}
	}

	return ifaces, nil
}

func (cd *CommandDNS) HasCommand(cmdName string) bool {
	_, err := exec.LookPath(cmdName)
	return err == nil
}
