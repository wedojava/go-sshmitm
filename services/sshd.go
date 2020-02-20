package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gliderlabs/ssh"
)

var User, Passwd, Local, Remote string

//var timestr = "./"+time.Now().Format("0102150405")+".txt"
var logFile = "2020.log"

func Server(privteKey, port string) {
	log.Println("starting ssh server on port " + port + "...")
	ssh.Handle(SessionHandler)
	log.Fatal(ssh.ListenAndServe(
		":"+port, nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			Passwd = pass
			return true
		}),
		ssh.HostKeyFile(privteKey)))
}

// SessionHandler will let session hold on forever
func SessionHandler(s ssh.Session) {
	User = s.User()
	Local = s.LocalAddr().String()
	Remote = s.RemoteAddr().String()
	content := fmt.Sprintf("%s\nUsername: %s\nPassword: %s\nLocal:%s\nRemote:%s\n",
		timeNow(), User, Passwd, Local, Remote)
	go logSave(content)
	log.Printf(content)
	time.Sleep(time.Duration(5) * time.Second)
}

func logSave(content string) {
	f, err := os.OpenFile("2020.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(content); err != nil {
		log.Println(err)
	}
}

func timeNow() string {
	loc := time.FixedZone("UTC+8", +8*60*60)
	modTime := time.Now()
	t := modTime.In(loc)
	return t.Format(time.RFC3339)
}
