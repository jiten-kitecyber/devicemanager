package dm

type IDNSDeviceManager interface {
	HasCommand(cmd string) bool
	SetDNS(iface string, primary string, secondary string) error
	GetActiveInterfaces() ([]string, error)
	GetDNS(string) (string, string, error)
	PostSetup() error
}

type IFirewallManager interface {
}
