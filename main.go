package main

import (
	"github.com/godbus/dbus"
	"log"
	"fmt"
)

func main() {

	conn, err := dbus.SystemBus()
	if err != nil {
		log.Fatalln(err)
	}
	conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"type='signal',path='/org/freedesktop/NetworkManager',member='StateChanged'")

	c := make(chan *dbus.Signal, 10)
	conn.Signal(c)
	for v := range c {
		fmt.Printf("%+v\n\n", v)
		// https://developer.gnome.org/NetworkManager/stable/nm-dbus-types.html#NMState
		if v.Body[0].(uint32) == 70 {
			log.Println("NM_STATE_CONNECTED_GLOBAL")
		} else if v.Body[0].(uint32) == 20 {
			log.Println("NM_STATE_DISCONNECTED")
		}
	}
}
