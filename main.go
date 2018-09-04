package main

import (
	"fmt"
	"image/jpeg"
	"math"
	"os"
)

func main() {
	fname := os.Args[1]
	imgf, err := os.Open(fname)

	if err != nil {
		fmt.Println("File not found!")
		os.Exit(1)
	}

	defer imgf.Close()

	img, err := jpeg.Decode(imgf)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Calculating...")

	rectangle := img.Bounds()

	var r, g, b float64 = 0, 0, 0
	w, h := rectangle.Max.X, rectangle.Max.Y
	i, j := 0, 0

	for i < w {
		for j < h {
			cr, cg, cb, _ := img.At(i, j).RGBA()

			r += math.Pow(float64(cr), 2)
			g += math.Pow(float64(cg), 2)
			b += math.Pow(float64(cb), 2)
			j++
		}
		j = 0
		i++
	}

	r = math.Sqrt(r / float64(w*h))
	g = math.Sqrt(g / float64(w*h))
	b = math.Sqrt(b / float64(w*h))

	fmt.Printf("#%02X%02X%02X\n", int(r/0x0101), int(g/0x0101), int(b/0x0101))
}
