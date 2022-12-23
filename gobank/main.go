package main

func main() {
	server := NewAPIServer(":3080")
	server.Run()
}
