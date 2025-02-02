package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/evlic/monkey"
)

func main() {
	monkey.Patch((*net.Dialer).DialContext, func(_ *net.Dialer, _ context.Context, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	}, monkey.OptGlobal)

	_, err := http.Get("http://taoshu.in")
	fmt.Println(err) // Get http://taoshu.in: no dialing allowed
}
