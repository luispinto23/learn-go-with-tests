package racer

import (
	"errors"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", errors.New("time out")
	}
}

func ping(url string) chan struct{} {
	// Why struct{} and not another type like a bool? Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
