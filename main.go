package main

import "introduction-gin/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
