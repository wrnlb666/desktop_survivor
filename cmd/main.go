package main

import (
    "os"
    "fmt"

    "survivor/cmd/menu"
    "survivor/cmd/server"
    "survivor/cmd/player"
    "survivor/cmd/object"
)

func main() {
    // parse cmd args
    args := os.Args[1:]
    if len(args) == 0 {
        // no args, run server
        fmt.Println("hello")
        server.Main()
    }

    // first argument should be appType
    switch args[0] {
    case server.AppType[server.PlayerType]:
        player.Main(args[1:])
    case server.AppType[server.ObjectType]:
        object.Main(args[1:])
    case server.AppType[server.MenuType]:
        menu.Main(args[1:])
    default:
        os.Exit(1)
    }
}
