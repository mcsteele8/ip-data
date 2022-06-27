package main

import (
	"ip-data/internal/spyware"
	"ip-data/tools/wlog"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	wlog.New().Info("IP Data Server Starting")
	spy := spyware.New(&http.Client{})
	lambda.Start(spy.GetSpywareInfo)
}
