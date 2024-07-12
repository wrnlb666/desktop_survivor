package server

import (
    rl  "github.com/gen2brain/raylib-go/raylib"
)

func getScreenSize() rl.Vector2 {
    rl.InitWindow(10, 10, "survivor")

    monitor := rl.GetCurrentMonitor()
    screenWidth := rl.GetMonitorWidth(monitor)
    screenHeight := rl.GetMonitorHeight(monitor)
    rl.CloseWindow()

    return rl.Vector2{
        X: float32(screenWidth),
        Y: float32(screenHeight),
    }
}

