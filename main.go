package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {

	reader := ExampleReadFrameAsJpeg("./test.mov", 5)
	img, err := imaging.Decode(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = imaging.Save(img, "./out2.jpeg")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
