package main

import hybrid_serverless "putrafirman.com/playground/hybrid-serverless"

func main() {
	e := hybrid_serverless.InitRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
