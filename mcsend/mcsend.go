package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("udp", "239.192.1.2:12345")
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    c := 0
    for {
        msg := fmt.Sprintf("msg%04d\n", c)
        conn.Write([]byte(msg))
        fmt.Println(c)
        time.Sleep(1 * time.Second)
        c++
    }
}
