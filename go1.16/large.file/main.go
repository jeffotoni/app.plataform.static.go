package main

import "embed"

func main() {
	//go:embed funcionario.json
	var f embed.FS
	data, _ := f.ReadFile("funcionario.json")
	print(string(data))
}
