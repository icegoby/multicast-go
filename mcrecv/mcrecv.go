package main

import (
    "fmt"
    "net"
)

func main() {
    udp_addr, err := net.ResolveUDPAddr("udp", "239.192.1.2:12345")
    if err != nil {
        panic(err)
    }
    iface, err := net.InterfaceByName("lo")
    if err != nil {
        panic(err)
    }
    listener, err := net.ListenMulticastUDP("udp", iface, udp_addr)
    if err != nil {
        panic(err)
    }
    defer listener.Close()

    buf := make([]byte, 1280)

    for {
        length, remote, err := listener.ReadFrom(buf)
        if err != nil {
            panic(err)
        }
        fmt.Printf("sender: %v\n", remote)
        fmt.Printf("Contents: %v\n", string(buf[:length]))
    }
}
