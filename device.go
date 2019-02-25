package wurfl

// Device object
type Device struct {
	Device              string
	Capabilities        KeyStoreList
	VirtualCapabilities KeyStoreList
}

// Release device object
func (d *Device) Release() {
	if d != nil {
		d.Capabilities.Release()
		d.VirtualCapabilities.Release()
	}
}
