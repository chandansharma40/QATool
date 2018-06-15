
package main

import "fmt"
import zmq "../"
import "time"
//import "net/http"
//import "html/template"

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)
	socket.Connect("tcp://192.168.43.250:5918")
	//socket.Connect("tcp://127.0.0.1:6000")

	for i := 0; i < 20; i++ {
		msg := fmt.Sprintf("{\"command\":\"getInterlocks\"}")
		socket.Send([]byte(msg), 0)
		println("Sending: ", msg)
		//socket.Recv(0)
		msg1, _ := socket.Recv(0)
		println("Got", string(msg1))
		time.Sleep(time.Second)
	}
}

/*package main

import (
    "log"
	"net"
	"strconv"
	"strings"
)

const (
	message       = "Ping"
	StopCharacter = "\r\n\r\n"
)

func SocketClient(ip string, port int) {
	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}

func main() {

	var (
		ip   = "127.0.0.1"
		port = 3333
	)

	SocketClient(ip, port)

}*/
