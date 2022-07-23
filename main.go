package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"syscall"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/pschlump/filelib"
	"github.com/pschlump/goqrcode"
	"github.com/pschlump/qr-secret/enc"
	"golang.org/x/term"
)

var encode = flag.String("encode", "", "file to encode")
var decode = flag.String("decode", "", "file to encode")
var output = flag.String("output", "", "file to encode")

var server = flag.Bool("server", false, "act as a webserver")
var hostPort = flag.String("host-port", "127.0.0.1:18410", "listen on host:port")

type EncHolder struct {
	Version  string `json:"v,omitempty"`
	Checksum string `json:"c,omitempty"`
	Part     string `json:"p,omitempty"`
	Of       string `json:"q,omitempty"`
	Data     string `json:"d"`
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "qr_gen_server: Usage: %s [flags]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse() // Parse CLI arguments to this, --cfg <name>.json

	fns := flag.Args()
	if len(fns) != 0 {
		fmt.Printf("Extra arguments are not supported [%s]\n", fns)
		os.Exit(1)
	}

	if *server && (*hostPort != "") {
		// do all setup to act as a server
	}

	// keyString := "Ya ya Ya ya" // xyzzy TODO - function to read in the password
	keyString, err := readPassword()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s on reading password\n", err)
		os.Exit(1)
	}

	out, err := filelib.Fopen(*output, "w")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to open %s for output: %s\n", *output, err)
		os.Exit(1)
	}
	defer out.Close()

	if *encode != "" {

		buf, err := ioutil.ReadFile(*encode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s on %s\n", err, *encode)
			os.Exit(1)
		}

		// xyzzy - TODO - put into format - add URL - urlBase64 encode ( 512 bytes? )
		// xyzzy - TODO - loop over chunks

		content := string(buf)
		redundancy := goqrcode.Highest
		size := 256

		encContent, err := enc.DataEncrypt([]byte(content), keyString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to encrypt %s Error: %s\n", *output, err)
			os.Exit(1)
		}

		// Generate the QR code in internal format
		var q *goqrcode.QRCode
		q, err = goqrcode.New(encContent, redundancy)
		goqrcode.CheckError(err)

		// Output QR Code as a PNG
		var png []byte
		png, err = q.PNG(size)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to create QR %s Error: %s\n", *output, err)
			os.Exit(1)
		}

		out.Write(png)

	} else if *decode != "" {

		encContent, err := DecodeQR(*decode)
		if err != nil {
			os.Exit(1)
		}
		content, err := enc.DataDecrypt(encContent, keyString)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to decrypt %s Error: %s\n", *output, err)
			os.Exit(1)
		}
		fmt.Fprintf(out, "%s", content)

	}
}

func DecodeQR(fn string) (data string, err error) {
	file, err := filelib.Fopen(fn, "r")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid file: %s error:%s\n", fn, err)
		return "", err
	}
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid QR code, file: %s error:%s\n", fn, err)
		return "", err
	}

	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to convert to QR-code to bitmap, file: %s error:%s\n", fn, err)
		return "", err
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to decode QR, file : %s error:%s\n", fn, err)
		return "", err
	}

	if db7 {
		fmt.Printf("%s: %s\n", fn, result)
	}
	return result.String(), nil
}

func readPassword() (password string, err error) {

	fmt.Print("Enter Password: ")
	if runtime.GOOS == "windows" {

		reader := bufio.NewReader(os.Stdin)
		password, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}

	} else {

		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return "", err
		}

		password = string(bytePassword)

	}

	return strings.TrimSpace(password), nil
}

const db7 = false

/* vim: set noai ts=4 sw=4: */
