package main

import (
	"encoding/gob"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"github.com/kataras/golog"
	"os"
	"time"
)

const MessageToSend = "Hey Palavian's\n\nReady to eat Fresh Fruits daily...?!\n\nWe are here to bring you Quality and Fresh Fruits Daily...\n\nOrder Now...\n\n• Anjeer - 100/Box\n\n• Apple - ₹120/Kg\n\n• Apple Royal Gala- ₹225/Kg\n\n• Mosambi - ₹140/kg\n\n• Plum - ₹250/Kg\n\n• Orange - ₹100/kg\n\n• Chiku- ₹80/kg\n\n• Anar- ₹120/kg\n\n• Custer Apple- ₹140/Kg\n\n• Pear(S.A)- ₹240/kg\n\n• Water Melon - ₹60 pc\n\n• Musk Melon - ₹60 Pc\n\n• Sweetcorn - ₹15/Pc\n\n• Litchi - ₹300/kg (Min 1 Kg)\n\n• Dates(Khaarek) - ₹100/Kg\n\n• Jamun - ₹300/Kg (Min 1Kg)\n\n• Elaichi Banana - ₹70/Dz\n\n• Dragon Fruit - ₹50/Pc\n\n• Kiwi - ₹50/Box(3pc)\n\n• Papaya - ₹40/Pc\n\n• Peach - ₹200/kg\n\n• Green Badam - ₹250/Kg\n\n• Kesar Mango - ₹550/ 5kg (Min 5 Kg)\n\n• Langda Mango - ₹450/ 5kg\n\n• Dasheri Mango - ₹450/ 5Kg\n\n\n\nPayment Mode :\nGpay/UPI Number -\n+91 9167908057\n\nContact Person -\nAshwin Gohil\n+91 9323430427\n\nEat Healthy, Stay Healthy...!\n🍏🍎🍐🍊🍌🍉🍇\U0001FAD0🍒🍑\U0001F96D\n\n**Free Home Delivery\n\nhttps://chat.whatsapp.com/JsqUz5Ej8yTGOEnkj1fOlp"
var contacts = []string{"919768187108", "919769908060"}

func main(){

	//=========New Connection
	wac, err := whatsapp.NewConn(30 * time.Second)
	if err != nil {
		panic(err)
	}
	wac.SetClientVersion(2, 2121, 6)

	err = login(wac)
	if err != nil {
		golog.Error(err)
	} else {
		golog.Info("login Successful")
	}

	<-time.After(3 * time.Second)

	//=========Looping through contacts and sending messages
	for i:=0;i<len(contacts);i++{

		img, err := os.Open("frusion.jpeg")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}

		msg := whatsapp.ImageMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: fmt.Sprintf("%s@s.whatsapp.net", contacts[i]),
			},
			Type:    "image/jpeg",
			Caption: MessageToSend,
			Content: img,
		}

		msgId, err := wac.Send(msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v", err)
			os.Exit(1)
		} else {
			fmt.Println("Message Sent -> ID : " + msgId)
		}

		//<-time.After(1 * time.Second)
		//fmt.Println("Sending Message to ",contacts[i])
		//
		//text := whatsapp.TextMessage{
		//	Info: whatsapp.MessageInfo{
		//		RemoteJid: fmt.Sprintf("%s@s.whatsapp.net", contacts[i]),
		//	},
		//	Text: MessageToSend,
		//}
		//
		//x, err := wac.Send(text)
		//fmt.Println(x)
		//if err!=nil {
		//	fmt.Println(err)
		//}
	}
}

func login(wac *whatsapp.Conn) error {

	session, err := readSession()
	if err == nil {
		session, err = wac.RestoreWithSession(session)
		if err == nil {
			return nil
		}
	}

	//==========New QR
	qr := make(chan string)
	go func() {
		terminal := qrcodeTerminal.New()
		terminal.Get(<-qr).Print()
	}()

	//=========Passing the QR scan to login
	session, err = wac.Login(qr)
	if err != nil {
		return fmt.Errorf("error during login: %v\n", err)
	}

	err = writeSession(session)
	if err!=nil {
	   	return fmt.Errorf("error saving session: %v\n", err)
	}

	return nil
}

func readSession() (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open("./whatsappSession.gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSession(session whatsapp.Session) error {
	file, err := os.Create("./whatsappSession.gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}


