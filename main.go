package main

import (
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const (
	AppName             = "GoCommunicator"
	publicKeyIdentifier = "gocom pk:"
	postUrlIdentifier   = "gocom target:"
)

var(
	uid , _= machineid.ProtectedID(AppName)
	hostname ,_= os.Hostname()
)

var (
	app = kingpin.New("jose-util", "A command-line utility for dealing with JOSE objects")

	// Upload message

	uploadMessageCommand = app.Command("msg", "Encrypt a plaintext, output ciphertext")

	uploadMessageText = uploadMessageCommand.Flag("text", "Path to input file (if applicable, stdin if missing)").Required().String()
	// uploadMessageTextTarget = uploadMessageCommand.Flag("target", "Path to input file (if applicable, stdin if missing)").Required().String()

	// Upload file

	uploadFileCommand = app.Command("file", "Encrypt a plaintext, output ciphertext")
	uploadFilePath    = uploadFileCommand.Flag("path", "Path to input file (if applicable, stdin if missing)").Required().String()
	// uploadFileTarget = uploadFileCommand.Flag("target", "Path to input file (if applicable, stdin if missing)").Required().String()

	target = app.Flag("target", "Path to input file (if applicable, stdin if missing)").Required().String()
)

func main() {
	app.Version("1.0.0.")
	app.UsageTemplate(kingpin.LongHelpTemplate)

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	switch command {
	case uploadMessageCommand.FullCommand():
		uploadMessage(*uploadMessageText, uid[:16],hostname )
	case uploadFileCommand.FullCommand():
		uploadFile()
	default:
		fmt.Fprintf(os.Stderr, "invalid command: %s\n", command)
		os.Exit(1)
	}
}
