package rtsmp

import (
    "log"
    "crypto/tls"
    "net"
    "bufio"
)

// setTLS provides optional parameter 
func setTLS(addr, pem, key string) {
	log.SetFlags(log.Lshortfile)

    cer, err := tls.LoadX509KeyPair(pem, key)
    if err != nil {
        log.Println(err)
        return
    }

    config := &tls.Config{Certificates: []tls.Certificate{cer}}
    ln, err := tls.Listen("tcp", addr, config) 
    if err != nil {
        log.Println(err)
        return
    }
    defer ln.Close()
}