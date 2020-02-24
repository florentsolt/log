package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/florentsolt/log"
)

func usage() {
	fmt.Printf("Usage: %s [hostname@ip:port]... OR - (for stdin)\n", os.Args[0])
	fmt.Println("Supported environment variables: FS_LOG_TCP FS_LOG_TCP_TIMEOUT FS_LOG_LEVEL FS_LOG_OUTPUT FS_LOG_NOCOLOR")
	os.Exit(1)
}

func main() {
	if len(os.Args) <= 1 {
		usage()
	}

	hosts, err := ParseHosts(os.Args[1:])
	if err != nil || hosts.IsEmpty() {
		fmt.Println(err)
		usage()
	}

	// timeout
	timeout, err := time.ParseDuration(os.Getenv("FS_LOG_TCP_TIMEOUT"))
	if err != nil && os.Getenv("FS_LOG_TCP_TIMEOUT") != "" {
		fmt.Println("Unable to pase timeout")
		os.Exit(2)
	}

	log.Console()

	// read stdin
	if hosts.IsStdin() {
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					return
				}
				fmt.Printf("Unable to read stdin: %s", err.Error())
				return
			}
			_, err = log.Write(line)
			if err != nil {
				// in case of error, only dump on stdin the line
				// fmt.Printf("Unable to write stderr: %s", err.Error())
				fmt.Printf("%s", line)
			}
		}
	} else {
		channel, err := hosts.Connect(timeout)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for {
			message := <-channel
			if len(message.Line) > 0 {
				log.Output().Write(message.HostnameWithPaddingAndSpace)
				_, err = log.Write(message.Line)
				if err != nil {
					// in case of error, only dump on stdin the line
					// fmt.Printf("Unable to write stderr: %s", err.Error())
					fmt.Printf("%s", message.Line)
				}
			}
		}
	}
}
