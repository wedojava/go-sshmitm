package services

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/gliderlabs/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

const TERMINAL = "powershell"

var Username, Password, LocalAddr, RemoteAddr string

func Server(privteKey, port string) {
	ssh.Handle(HoldOnSessionHandler)
	log.Println("starting ssh server on port " + port + "...")
	log.Fatal(ssh.ListenAndServe(
		":"+port, nil,
		ssh.PasswordAuth(func(ctx ssh.Context, pass string) bool {
			Password = pass
			return true
		}),
		ssh.HostKeyFile(privteKey)))
}

// sessionHandler handling established SSH sessions
func sessionHandler(s ssh.Session) {
	_, _, isPty := s.Pty()
	if isPty {
		fmt.Println("PTY requested")

		cmd := exec.Command(TERMINAL)
		stdin, err := cmd.StdinPipe()
		if err != nil {
			panic(err)
		}
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			panic(err)
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			panic(err)
		}

		go func() {
			Username = s.User()
			LocalAddr = s.LocalAddr().String()
			RemoteAddr = s.RemoteAddr().String()
			io.Copy(stdin, s)
		}()
		go func() {
			io.Copy(s, stdout)
		}()
		go func() {
			io.Copy(s, stderr)
		}()

		err = cmd.Run()
		if err == nil {
			log.Println("session ended normally")
			s.Exit(0)
		} else {
			log.Printf("session ended with an error: %v\n", err)

			exitCode := 1
			if exitError, ok := err.(*exec.ExitError); ok {
				exitCode = exitError.ExitCode()
				log.Printf("exit code: %d\n", exitCode)
			}

			s.Exit(exitCode)
		}
	} else {
		io.WriteString(s, "No PTY requested.\n")
		s.Exit(1)
	}
}

// HoldOnSessionHandler will let session hold on forever
func HoldOnSessionHandler(s ssh.Session) {
	Username = s.User()
	LocalAddr = s.LocalAddr().String()
	RemoteAddr = s.RemoteAddr().String()
	term := terminal.NewTerminal(s, "")
	line := ""
	go func() {
		content := []byte("Username: " + Username + "\nPassword: " + Password + "\nLocalAddr: " + LocalAddr + "\nRemoteAddr: " + RemoteAddr)
		_ = ioutil.WriteFile("./"+time.Now().Format("0102150405")+".txt", content, os.ModePerm)
	}()
	for {
		line, _ = term.ReadLine()
		if line == "" || line == " " {
			time.Sleep(time.Duration(5) * time.Second)
			break
		}
	}
}

// echoSH is sessionHandler to echo client's terminal
func echoSH(s ssh.Session) {
	term := terminal.NewTerminal(s, "> ")
	line := ""
	for {
		line, _ = term.ReadLine()
		if line == "quit" || line == "exit" {
			break
		}
		io.WriteString(s, fmt.Sprintf("You wrote: %s\n", line))
	}
}
