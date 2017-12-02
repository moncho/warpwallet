package qr

import (
	"io"
	"io/ioutil"

	"rsc.io/qr"
)

const blackSquare = "\033[40m  \033[0m"
const whiteSquare = "\033[47m  \033[0m"

var newLine = []byte("\n")

type encodingParams struct {
	Level     qr.Level
	Reader    io.Reader
	Writer    io.Writer
	BlackChar string
	WhiteChar string
}

//Creates a QR that can be displayed on a terminal
func encode(params encodingParams) error {
	w := params.Writer
	white := params.WhiteChar
	black := params.BlackChar
	var text string

	if b, err := ioutil.ReadAll(params.Reader); err == nil {
		text = string(b)
	} else {
		return err
	}
	code, err := qr.Encode(text, params.Level)

	if err != nil {
		return err
	}

	textVersion := qrToText(code, []byte(white), []byte(black))
	_, err = w.Write(textVersion)
	return err
}

func qrToText(code *qr.Code, white, black []byte) []byte {

	textVersion := white
	for i := 0; i <= code.Size; i++ {
		textVersion = append(textVersion, white...)
	}
	textVersion = append(textVersion, newLine...)

	for i := 0; i <= code.Size; i++ {
		textVersion = append(textVersion, white...)
		for j := 0; j <= code.Size; j++ {
			if code.Black(i, j) {
				textVersion = append(textVersion, black...)
			} else {
				textVersion = append(textVersion, white...)
			}
		}
		textVersion = append(textVersion, newLine...)
	}

	textVersion = append(textVersion, newLine...)
	return textVersion
}

//Copy encodes the given reader as a QR code in text mode and copies it
//to the given writer
func Copy(w io.Writer, r io.Reader) error {

	//func Encode(text string, w io.Writer) error {
	params := encodingParams{
		Level:     qr.L,
		Reader:    r,
		Writer:    w,
		BlackChar: blackSquare,
		WhiteChar: whiteSquare,
	}
	return encode(params)
}
