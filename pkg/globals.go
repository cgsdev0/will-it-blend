package pkg

import (
	"net/http"
	"time"
)

var CLIENT_ID = "o1ffndtnzdf5xa6efbaytdti617btp"

var httpClient = http.Client{
	Timeout: time.Second * 10,
}
