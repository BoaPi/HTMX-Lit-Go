package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func Timer(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	currentSeconds := currentTime.Second()

	fmt.Fprintf(w, "Current seconds: %s", strconv.Itoa(currentSeconds))
}
