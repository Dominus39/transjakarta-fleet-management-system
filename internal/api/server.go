package api

func StartServer(port string) {
	r := SetupRouter()
	if err := r.Run(port); err != nil {
		panic("Failed to start the server: " + err.Error())
	}
}
