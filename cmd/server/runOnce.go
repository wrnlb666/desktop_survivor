package server

import (
    "os"
    
    "survivor/util"

    rl  "github.com/gen2brain/raylib-go/raylib"
)

func runOnce() {
    if util.DBExist() {
        // db exists, game already running
        // create window to tell player if it is not desired
        const width = 800
        const height = 600
        rl.InitWindow(width, height, "survivor")

        monitor := rl.GetCurrentMonitor()
        screenWidth := rl.GetMonitorWidth(monitor)
        screenHeight := rl.GetMonitorHeight(monitor)
        rl.SetWindowPosition(
            (screenWidth - width) / 2,
            (screenHeight - height) / 2)
        rl.SetTargetFPS(30)

        for !rl.WindowShouldClose() {
            rl.BeginDrawing()
            rl.ClearBackground(rl.Black)
            rl.DrawText("Game already running", 10, 10, 64, rl.RayWhite)
            rl.DrawText("Try deleting", 10, 110, 64, rl.RayWhite)
            rl.DrawText("`desktop_survivor.db`", 10, 210, 64, rl.RayWhite)
            rl.DrawText("if it is not desired", 10, 310, 64, rl.RayWhite)
            rl.EndDrawing()
        }
        rl.CloseWindow()
        os.Exit(1)
    }
}


