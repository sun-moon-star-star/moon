package text

import (
    "fmt"
    "math/rand"
    "time"
    "os"
)

func RandSeedTime() {
    rand.Seed(time.Now().Unix()^ int64(os.Getpid()))
}

func GetUUID() (string, error) {
    b := make([]byte, 16)
    _, err := rand.Read(b)
    if err != nil {
        return "", err
    }
    uuid := fmt.Sprintf("%x%x%x%x%x",
        b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
    return uuid, nil
}