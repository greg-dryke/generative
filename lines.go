// from https://yourbasic.org/golang/create-image/

package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
    "fmt"
    "math/rand"
    "flag"
)

var debug bool = false

func draw(img *image.RGBA, c color.Color, x int, y int, xstep int, ystep int) {
    leftToRight := rand.Intn(2)
    for i := 0; i < xstep; i++ {
        if leftToRight > 0 {
            img.Set(x+i, y+i, c)
        } else {
            img.Set(x+xstep-i, y+i, c)
        }
    }
}


func main() {
    var width = flag.Int("width", 200, "width of the image")
    var height = flag.Int("height", 100, "height of the image")
    var step = flag.Int("step", 20, "step for lines")
    var outputFile = flag.String("output", "output.png", "name of the output png")
    flag.Parse()
    // width := 1080
    // height := 2160

    upLeft := image.Point{0, 0}
    lowRight := image.Point{*width, *height}

    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    // Colors are defined by Red, Green, Blue, Alpha uint8 values.
    cyan := color.RGBA{100, 200, 200, 0xff}

    /*
         idea from: https://generativeartistry.com/tutorials/tiled-lines/
    */

    for x := 0; x < *width; x+=*step {
        for y := 0; y < *height; y+=*step {
            if debug {
                fmt.Println("x:", x, "y:", y)
            }
            draw(img, cyan, x,y,*step,*step)
        }
    }

    // Encode as PNG.
    f, _ := os.Create(*outputFile)
    png.Encode(f, img)
}
