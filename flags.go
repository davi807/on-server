package main

import (
	"errors"
	"flag"
	"net"
	"os"
)

var flags struct {
	ip      string
	port    string
	showIP6 bool

	path        string
	messagePath string
	uploadPath  string

	noMessage bool
	noUpload  bool
	noFiles   bool
}

func initFlags() error {
	flag.StringVar(&flags.ip, "ip", "", "Server ip address, required IPv4, if not provided listens to all interfaces")
	flag.StringVar(&flags.port, "port", "2100", "Port number")
	flag.BoolVar(&flags.showIP6, "show-ip6", false, "Show IP6 addresses in list if 'ip' not provided")

	flag.StringVar(&flags.path, "path", "", "Server files root (default current path)")
	flag.StringVar(&flags.messagePath, "message-path", "", "Text message files location (default current path)")
	flag.StringVar(&flags.uploadPath, "upload-path", "", "Uploaded files location (default current path)")

	flag.BoolVar(&flags.noMessage, "no-message", false, "Disable text submit")
	flag.BoolVar(&flags.noUpload, "no-upload", false, "Disable file upload")
	flag.BoolVar(&flags.noFiles, "no-files", false, "Disable files listing")

	flag.Parse()

	return validateFlags()
}

func validateFlags() error {
	// is address avaluble
	if len(flags.ip) > 0 {
		c, err := net.Listen("tcp", flags.ip+":"+flags.port)
		if err != nil {
			return errors.New("Cant bind to address \n" + err.Error())
		}
		c.Close()
	}

	// get working direcctory
	var root string
	root, e := os.Getwd()
	if e != nil {
		return errors.New("Can't get working directory path \n" + e.Error())
	}

	// server paths
	if len(flags.path) > 0 {
		if info, err := os.Stat(flags.path); err != nil || !info.IsDir() {
			return errors.New("Wrong 'path' argument")
		}
	} else {
		flags.path = root
	}

	if len(flags.messagePath) > 0 {
		if info, err := os.Stat(flags.messagePath); !(err == nil && info.IsDir()) {
			return errors.New("Wrong 'message-path' argument")
		}
	} else {
		flags.messagePath = flags.path
	}

	if len(flags.uploadPath) > 0 {
		if info, err := os.Stat(flags.uploadPath); !(err == nil && info.IsDir()) {
			return errors.New("Wrong 'upload-path' argument")
		}
	} else {
		flags.uploadPath = flags.path
	}

	// if everything is disabled
	if flags.noMessage && flags.noUpload && flags.noFiles {
		return errors.New("All server functions is disabled")
	}

	return nil
}