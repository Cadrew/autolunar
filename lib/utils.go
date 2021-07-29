package autolunar

import (
	"time"
)

func contains(slice []int, item int) bool {
    set := make(map[int]struct{}, len(slice))
    for _, s := range slice {
        set[s] = struct{}{}
    }
    _, ok := set[item] 
    return ok
}

func getTimestamp() int64 {
    return time.Now().UnixNano() / int64(time.Millisecond)
}