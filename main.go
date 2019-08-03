package main

import (
	"flag"
	"image/png"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/skip2/go-qrcode"
)

const (
	// QrSizeW -
	QrSizeW = 200
	// QrSizeH -
	QrSizeH = 200
	// FileName1 -
	FileName1 = "qrcode1.png"
	// FileName2 -
	FileName2 = "qrcode2.png"
	// URL -
	URL = "https://www.yahoo.co.jp/"
)

var url = flag.String("url", "", "URL of QR code conversion source")

func main() {
	flag.Parse()

	createQr1(*url)

	// ↓こっちの方がいい
	createQr2(*url)
}

//"github.com/boombuler/barcode/qr"
func createQr1(url string) {
	// Create the barcode
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

	// Scale the barcode to 200x200 pixels
	qrCode, _ = barcode.Scale(qrCode, QrSizeW, QrSizeH)

	// create the output file
	file, _ := os.Create(FileName1)
	defer file.Close()

	// encode the barcode as png
	png.Encode(file, qrCode)
}

//"github.com/skip2/go-qrcode"
func createQr2(url string) {
	if err := qrcode.WriteFile(url, qrcode.Medium, 256, FileName2); err != nil {
		panic(err)
	}
}
