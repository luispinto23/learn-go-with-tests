package racer

import (
	"errors"
	"net/http"
	"time"
)

// SELECT
// - Helps you wait on multiple channels.
// - Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

// httptest
// - A convenient way of creating test servers so you can have reliable and controllable tests.
// - Uses the same interfaces as the "real" net/http servers which is consistent and less for you to learn.

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
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
