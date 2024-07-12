package server


import (

)


type appType int
const (
    ServerType appType = iota
    PlayerType
    ObjectType
    MenuType
)
var AppType map[appType]string = map[appType]string{
    ServerType: "server",
    PlayerType: "player",
    ObjectType: "object",
    MenuType:   "menu",
}
