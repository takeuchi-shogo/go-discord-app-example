package main

import api "github.com/takeuchi-shogo/go-discord-app/backend/service/api/runner"

func main() {
	api.NewApiServer()
	api.NewWorker()
}
