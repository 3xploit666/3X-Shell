package main

import (
	"3X-Shell/3x-Shell/help"
)

func main() {
	info := "192.168.1.9:9090"
	help.Conexion3x("tcp", info)
}
