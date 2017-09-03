package qr

import (
	"io"

	"rsc.io/qr"
)

const blackSquare = "\033[40m  \033[0m"
const whiteSquare = "\033[47m  \033[0m"

type encodingParams struct {
	Level     qr.Level
	Writer    io.Writer
	BlackChar string
	WhiteChar string
}

//Creates a QR that can be displayed on a terminal
func encode(text string, params encodingParams) error {
	w := params.Writer
	white := params.WhiteChar
	black := params.BlackChar
	code, err := qr.Encode(text, params.Level)

	if err != nil {
		return err
	}

	w.Write([]byte(white))
	for i := 0; i <= code.Size; i++ {
		w.Write([]byte(white))
	}
	w.Write([]byte("\n"))

	for i := 0; i <= code.Size; i++ {
		w.Write([]byte(white))
		for j := 0; j <= code.Size; j++ {
			if code.Black(i, j) {
				w.Write([]byte(black))
			} else {
				w.Write([]byte(white))
			}
		}
		w.Write([]byte("\n"))
	}
	return err
}

//Encode encodes the given text as a QR code, the QR is written
//on the given writer
func Encode(text string, w io.Writer) error {
	params := encodingParams{
		Level:     qr.L,
		Writer:    w,
		BlackChar: blackSquare,
		WhiteChar: whiteSquare,
	}
	return encode(text, params)
}
