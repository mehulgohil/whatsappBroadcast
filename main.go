package main

import (
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"os"
	"time"
)

var contacts = []string{"919167908060", "918080908057", "919769908060", "919324498562", "919768187108"}

func main(){

	wac, err := whatsapp.NewConn(30 * time.Second)
	if err != nil {
		panic(err)
	}
	wac.SetClientVersion(2, 2121, 6)

	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	session, err := wac.Login(qr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during login: %v\n", err)
	}
	fmt.Printf("login successful, session: %v\n", session)

	for i:=0;i<len(contacts);i++{
		fmt.Println("Sending Message to ",contacts[i])

		text := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: fmt.Sprintf("%s@s.whatsapp.net", contacts[i]),
			},
			Text: "Hello Whatsapp, Testing",
		}

		x, err := wac.Send(text)
		fmt.Println(x)
		if err!=nil {
			fmt.Println(err)
		}
	}

	time.Sleep(1000 * time.Second)

}
