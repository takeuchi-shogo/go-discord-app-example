package main

import api "github.com/takeuchi-shogo/go-worker-scheduler-app-example/backend/internal/api/runner"

func main() {
	api.NewApiServer()
	api.NewWorker()
}
