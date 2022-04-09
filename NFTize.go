package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"time"
)

func errHandle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	for gen := 0; gen < 50; gen++ {
		time.Sleep(5)
		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)
		seedr := r.Float64() * 1.5
		seedg := r.Float64() * 1.5
		seedb := r.Float64() * 1.5
		fmt.Println(seedr)
		file, err := os.Open("in.jpg")
		errHandle(err)
		in, err := jpeg.Decode(file)
		errHandle(err)
		b := in.Bounds()
		m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
		draw.Draw(m, m.Bounds(), in, b.Min, draw.Src)
		for i := 0; i < in.Bounds().Max.X; i++ {
			for j := 0; j < in.Bounds().Max.Y; j++ {
				r, g, b, a := in.At(i, j).RGBA()
				m.Set(i, j, color.RGBA{uint8(int(float64(r)*seedr) / 257), uint8(int(float64(g)*seedg) / 257), uint8(int(float64(b)*seedb) / 257), uint8(a / 257)})
			}
		}
		str := fmt.Sprintf("out/out-%v.jpg", gen)
		os.Mkdir("out", 0755)
		out, err := os.Create(str)
		defer out.Close()
		errHandle(err)
		jpeg.Encode(out, m, nil)
	}
}