package util


import (
    "time"
)


func Must(err error) {
    if err != nil {
        panic(err)
    }
}


func GetDeltaTimeFunc() func() float32 {
    last := time.Now()
    return func() float32 {
        curr := time.Now()
        delta := time.Duration(curr.Sub(last)).Seconds()
        last = curr
        return float32(delta)
    }
}

