package main

import (
    "fmt"
    "flag"
)

type Config struct {
    Verbose bool
    Debug   bool
    Host    string
    Port    int
}

func main () {
    cfg := parseFlags()

    fmt.Println("Verbose:", cfg.Verbose)
    fmt.Println("Debug:", cfg.Verbose)
    fmt.Println("Host:", cfg.Host)
    fmt.Println("Port:", cfg.Port)
}

func parseFlags() Config {
    var cfg Config
    flag.BoolVar(&cfg.Verbose, "verbose", true, "Verbose mode")
    flag.BoolVar(&cfg.Verbose, "debug", false, "Debug mode")
    flag.StringVar(&cfg.Host, "host", "localhost", "Host to run on")
    flag.IntVar(&cfg.Port, "port", 12345, "Port to run on")
    flag.Parse()
    return cfg
}
