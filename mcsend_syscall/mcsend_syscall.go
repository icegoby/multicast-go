package main

import (
    // "fmt"
    "syscall"
)

func main() {
    sin := &syscall.SockaddrInet4 {
        Port: 12345,
        Addr: [4]byte{239, 192, 1, 2},
    }

    src := &syscall.SockaddrInet4 {
        Port: 22345,
        Addr: [4]byte{127, 0, 0, 1},
    }

    sock, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
    if err != nil {
        panic(err)
    }
    defer syscall.Close(sock)

    err = syscall.Bind(sock, src)
    if err != nil {
        panic(err)
    }
    err = syscall.SetsockoptByte(sock, syscall.IPPROTO_IP, syscall.IP_MULTICAST_LOOP, 1)
    if err != nil {
        panic(err)
    }
    msg := "hello"
    err = syscall.Sendto(sock, []byte(msg), 0, sin)
    if err != nil {
        panic(err)
    }
}
