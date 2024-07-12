package util


import (
    "math"

    win "github.com/lxn/win"
    _   "github.com/glebarez/go-sqlite"
    rl  "github.com/gen2brain/raylib-go/raylib"
)


func Normalize(dst, src rl.Vector2) rl.Vector2 {
    disX := dst.X - src.X
    disY := dst.Y - src.Y
    if disX == 0 && disY == 0 {
        return rl.Vector2{
            X: 0,
            Y: 0,
        }
    }
    absSum := float32(math.Abs(float64(disX)) + math.Abs(float64(disY)))
    return rl.Vector2{
        X: disX / absSum,
        Y: disY / absSum,
    }
}


func GetCursorPos() rl.Vector2 {
    var pos win.POINT
    win.GetCursorPos(&pos)
    return rl.Vector2{
        X: float32(pos.X),
        Y: float32(pos.Y),
    }
}

