package main

import (
	"image/color"
	"io"
	"math/rand"
	"image/gif"
	"image"
	"math"
	"net/http"
	"log"
	"strconv"
)

var palette = []color.Color{color.Black, color.RGBA{0x00, 0xFF, 0x00, 0xff }, color.RGBA{0xff, 0x00,0x00,0xff}, color.RGBA{0x00, 0x00, 0xff, 0xff}}

const(
	blackindex = 0
	greenindex = 1
	redindex = 2
	blueindex = 3
)

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	cycles, _ := strconv.Atoi(r.URL.Query()["cycles"][0])
	lissajous(w, float64(cycles))
}

func lissajous(out io.Writer, cycles float64){
	const(
		//cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() *3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++{
		rect := image.Rect(0,0,2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t+=res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(10)%4))
		}
		phase+=0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
