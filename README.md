# DESKTOP SURVIVOR

A [vampire survivors](https://store.steampowered.com/app/1794680/Vampire_Survivors/) clone using [`golang`](https://go.dev/), [`raylib`](https://github.com/raysan5/raylib), and [`raylib-go`](https://github.com/gen2brain/raylib-go). 

## Caution
- Due to the usage of various windows APIs, this project is not cross-platform, and currently is only available for windows. 
- As `cgo` is slow, and windows is the only platform that allows us to avoid using cgo even when calling C function(which is necessary for using `raylib`, `sqlite`). This game will probably not be ported to other platform. 
- Since this project is using [`raylib-go`](https://github.com/gen2brain/raylib-go), checkout [this](https://github.com/gen2brain/raylib-go?tab=readme-ov-file#purego-without-cgo-ie-cgo_enabled0) to install dependency. 
