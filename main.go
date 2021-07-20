package main

var config_version_changed map[string]bool

func main() {
	r := InitHttpServerrouter()
	r.Run(":8050")
}
