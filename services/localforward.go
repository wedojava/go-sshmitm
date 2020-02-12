package services

import (
	"fmt"
	"github.com/ScaleFT/sshkeys"
	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
)

const PORT_LOCAL = "2222"

func LocalForward() {
	log.Println("starting ssh server on port " + PORT_LOCAL + "...")

	server := ssh.Server{
		LocalPortForwardingCallback: ssh.LocalPortForwardingCallback(func(ctx ssh.Context, dhost string, dport uint32) bool {
			dhost = "127.0.0.1"
			dport = 80
			log.Println("Accepted forward", dhost, dport)
			return true
		}),
		Addr: ":" + PORT_LOCAL,
		Handler: ssh.Handler(HoldOnSessionHandler),
		ChannelHandlers: map[string]ssh.ChannelHandler{
			"session":      ssh.DefaultSessionHandler,
			"direct-tcpip": ssh.DirectTCPIPHandler,
		},
		PasswordHandler: func(ctx ssh.Context, pass string) bool {
			fmt.Println("Password: ", pass)
			return true
		},
	}
	s, _ := getKeyFile("./id_rsa", "")
	server.AddHostKey(s) // PublicKey

	log.Fatal(server.ListenAndServe())
}

// returns ssh.Signer from user you running app home path + cutted key path.
// (ex. pubkey,err := getKeyFile("/.ssh/id_rsa") )
func getKeyFile(keypath, passphrase string) (ssh.Signer, error) {
	var pubkey gossh.Signer
	var err error
	buf, err := ioutil.ReadFile(keypath)
	if err != nil {
		return nil, err
	}

	if passphrase != "" {
		pubkey, err = sshkeys.ParseEncryptedPrivateKey(buf, []byte(passphrase))
	} else {
		pubkey, err = gossh.ParsePrivateKey(buf)
	}

	if err != nil {
		return nil, err
	}

	return pubkey, nil
}
