package html2pdfsocket

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/mailru/easyjson"
)

type (
	ForPdf struct {
		Url  string
		Html string
	}

	Connect struct {
		Address string
		Port    string
	}
)

func GetPdf(connect Connect, params ForPdf) []byte {
	addr := strings.Join([]string{connect.Address, connect.Port}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	out, _ := easyjson.Marshal(params)
	conn.Write([]byte(out))

	buf := []byte{}          // big buffer
	tmp := make([]byte, 256) // using small tmo buffer for demonstrating
	for {
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		buf = append(buf, tmp[:n]...)

	}
	return buf
}
