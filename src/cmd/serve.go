package cmd

import (
	"encoding/binary"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
)

const SocketFile = "lguf.sock"

const (
	MSG_TYPE_GET byte = iota
	MSG_TYPE_SET
	MSG_TYPE_ERROR
)

/*
  This server allows us to get around permission issues. The normal CLI needs
  root access in order to do it's job. We get setuid the binary in a normal
  desktop environment, but with Snap that doesn't work. Instead, we run a
  daemon to pass the user run CLI commands into a root run one.

  This is a clever trick recommended in the Snapcraft Forum:
  https://forum.snapcraft.io/t/how-to-add-or-workaround-a-udev-rule/12829/2?u=felicianotech

  Outside of snaps, we can just use a udev rule instead.
*/
func bridgeServer(c net.Conn) {

	data := make([]byte, 3)

	log.Printf("Client connected [%s]", c.RemoteAddr().Network())
	size, err := c.Read(data)
	if err != nil {
		log.Fatal("Fatal: lguf-cli (daemon) Couldn't read from socket.")
	}
	defer c.Close()

	if size != 3 {
		log.Print("Error: Received data with incorrect size.")
	}

	//DEBUG
	// Below is a debug implementation. Will return the value 1080 as a test.
	// Next up is to make this return a real brightness, or set one.
	// int 1080 -> 0x0438
	fakeBrightness := make([]byte, 2)
	size = binary.PutUvarint(fakeBrightness, 0x0438)
	if size != 2 {
		log.Fatal("Writing fake brightness wasn't 2 bytes.")
	}

	data[0] = MSG_TYPE_GET
	data[1] = fakeBrightness[0]
	data[2] = fakeBrightness[1]

	size, err = c.Write(data)
	if err != nil {
		log.Fatal("Fatal: lguf-cli (daemon) Couldn't read from socket.")
	}
	if size != 3 {
		log.Print("Error: Wrote data, incorrect size returned.")
	}
}

var serveCmd = &cobra.Command{
	Use:    "serve",
	Short:  "Run this program in daemon mode",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {

		socketPath := os.Getenv("SNAP_COMMON") + "/" + SocketFile

		if err := os.RemoveAll(socketPath); err != nil {
			log.Fatal(err)
		}

		listener, err := net.Listen("unix", socketPath)
		if err != nil {
			log.Fatal("Listen error:", err)
		}
		defer listener.Close()

		if err := os.Chmod(socketPath, 0666); err != nil {
			log.Fatal(err)
		}

		for {

			// As new connections come in, send them to bridgeServer in a
			// goroutine.
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}

			go bridgeServer(conn)
		}
	},
}

func init() {

	rootCmd.AddCommand(serveCmd)
}
