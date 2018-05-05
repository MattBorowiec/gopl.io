// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())
	config := map[string]float64{
		"cycles":  5,     // number of complete x oscillator revolutions
		"res":     0.001, // angular resolution
		"size":    100,   // image canvas covers [-size..+size]
		"nframes": 64,    // number of animation frames
		"delay":   8,     // delay between frames in 10ms units
	}

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				fmt.Println("Error parsing url", err.Error())
			}

			//ERROR HANDLING LOL
			cycles, _ := strconv.ParseFloat(r.Form["cycles"][0], 64)
			config["cycles"] = cycles
			res, _ := strconv.ParseFloat(r.Form["res"][0], 64)
			config["res"] = res
			size, _ := strconv.ParseFloat(r.Form["size"][0], 64)
			config["size"] = size
			nframes, _ := strconv.ParseFloat(r.Form["nframes"][0], 64)
			config["nframes"] = nframes
			delay, _ := strconv.ParseFloat(r.Form["delay"][0], 64)
			config["delay"] = delay
			lissajous(w, config)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	//lissajous(os.Stdout, config)
}

func lissajous(out io.Writer, config map[string]float64) {

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: int(config["nframes"])}
	phase := 0.0 // phase difference
	for i := 0; i < int(config["nframes"]); i++ {
		rect := image.Rect(0, 0, 2*int(config["size"])+1, 2*int(config["size"])+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < config["cycles"]*2*math.Pi; t += config["res"] {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(config["size"])+int(x*config["size"]+0.5), int(config["size"])+int(y*config["size"]+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(config["delay"]))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
