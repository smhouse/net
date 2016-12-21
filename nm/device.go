package nm

import (
	"github.com/godbus/dbus"
	"log"
	"fmt"
	"math"
)



type Device_t struct {
	Path		dbus.ObjectPath
	IPv4		string
	Type		string
}

var types = map[uint32]string {
	0:	"Unknown",
	14:	"Generic",
	1:	"Ethernet",
	2:	"WiFi",
	5:	"Bluetooth",
	6:	"OLPC_MESH",
	7:	"WIMAX",
	8:	"MODEM",
	9:	"INFINIBAND",
	10:	"BOND",
	11:	"VLAN",
	12:	"ADSL",
	13:	"BRIDGE",
	15:	"TEAM",
	16:	"TUN",
	17:	"TUNNEL",
	18:	"MACVLAN",
	19:	"VXLAN",
	20:	"VETH",
}

func convertType(t uint32) string {
	if val, ok := types[t]; ok {
		return val
	}

	return "Unknown"
}

func convertIP(decimal uint32) string {
	arr := []float64{0, 0, 0, 0}
	c := 16777216.0
	ip := float64(decimal)

	for i := 0; i < 4; i++ {
		k := math.Floor(ip / c)
		ip -= c * k
		arr[i] = k
		c /= 256.0
	}

	return fmt.Sprintf("%.0f.%.0f.%.0f.%.0f", arr[3], arr[2], arr[1], arr[0])
}

func GetDevices() (*[]Device_t, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.NetworkManager", "/org/freedesktop/NetworkManager")
	var devices []dbus.ObjectPath
	err = obj.Call("org.freedesktop.NetworkManager.GetDevices", 0).Store(&devices)
	if err != nil {
		return nil, err
	}

	result := make([]Device_t, 0)
	for _, v := range devices {
		deviceObject := conn.Object("org.freedesktop.NetworkManager", v)
		variant, err := deviceObject.GetProperty("org.freedesktop.NetworkManager.Device.Ip4Address")
		if err != nil {
			log.Fatalln(err)
			return nil, nil
		}
		_ipv4 := convertIP(variant.Value().(uint32))

		variant, err = deviceObject.GetProperty("org.freedesktop.NetworkManager.Device.DeviceType")
		if err != nil {
			log.Fatalln(err)
			return nil, nil
		}
		_type := convertType(variant.Value().(uint32))

		d := Device_t{
			Path:	v,
			IPv4:	_ipv4,
			Type:	_type,
		}

		result = append(result, d)
	}

	return &result, nil
}