package nm

import (
	"github.com/godbus/dbus"
)

func GetDevices() (*[]dbus.ObjectPath, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	var s interface{}
	err = obj.Call("org.freedesktop.NetworkManager.GetDevices", 0).Store(&s)
	if err != nil {
		return nil, err
	}

	ret := s.([]dbus.ObjectPath)
	return &ret, nil
}