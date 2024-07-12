package player


import (
	// "survivor/ui"
	"survivor/util"

    // rl  "github.com/gen2brain/raylib-go/raylib"
)


var status Status


func start() {
    setFlags()
    _, err := util.DBConnect()
    util.Must(err)

    // read global
    query := util.DBQuery()
    status.Global, err = query.GetGlobal(ctx)
    util.Must(err)
}

