package server

import (
	"fmt"
	"time"
)

const sleepTime = 1

func Hello() string {
	time.Sleep(sleepTime * time.Second)
	return fmt.Sprintf("[%s] Hello. I am server.", time.Now())
}
