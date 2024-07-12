package util


import (
    "survivor/ui"

    rl  "github.com/gen2brain/raylib-go/raylib"
)


func GetScreenSize(global ui.Global) rl.Vector2 {
    return rl.Vector2{
        X: float32(global.ScreenWidth),
        Y: float32(global.ScreenHeight),
    }
}
