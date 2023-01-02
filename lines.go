// from https://yourbasic.org/golang/create-image/

package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
    "fmt"
    "flag"
    "strings"
        crypto_rand "crypto/rand"
    "encoding/binary"
    math_rand "math/rand"
)

var debug bool = false

func draw(img *image.RGBA, c color.Color, x int, y int, xstep int, ystep int, width int) {
    leftToRight := math_rand.Intn(2)
    fmt.Println("x", x)
    fmt.Println("random num:", leftToRight)
    xVal := 0
    for i := 0; i < xstep; i++ {
        if leftToRight > 0 {
            xVal = x+i
        } else {
            xVal = x+xstep-i
        }
        for j := 0; j < width; j++ {
            img.Set(xVal, y+i+j, c)
            img.Set(xVal, y+i-j, c)
        }
    }
}

// https://stackoverflow.com/questions/12321133/how-to-properly-seed-random-number-generator
func init() {
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func main() {
    var width = flag.Int("width", 200, "width of the image")
    var height = flag.Int("height", 100, "height of the image")
    var step = flag.Int("step", 20, "step for lines")
    var lineWidth = flag.Int("lwidth", 3, "width for lines")
    var outputFile = flag.String("output", "output.png", "name of the output png")
    flag.Parse()
    // width := 1080
    // height := 2160


    upLeft := image.Point{0, 0}
    lowRight := image.Point{*width, *height}

    img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

    // Colors are defined by Red, Green, Blue, Alpha uint8 values.
    //cyan := color.RGBA{100, 200, 200, 0xff}
    //purple := color.RGBA{200, 200, 100, 0xff}
    //yellow := color.RGBA{200, 100, 200, 0xff}
    // https://twitter.com/colorschemez/status/1406467329827885057
    lightBlue := color.RGBA{162,207,254, 0xff}
    brightGreen := color.RGBA{174,255,110, 0xff}
    midGray := color.RGBA{146,149,145,0xff} 
    var colors [3]color.Color
    colors[0] = lightBlue
    colors[1] = midGray
    colors[2] = brightGreen
    /*
         idea from: https://generativeartistry.com/tutorials/tiled-lines/
    */

    for x := 0; x < *width; x+=*step {
        for y := 0; y < *height; y+=*step {
            c := colors[math_rand.Intn(3)]
            if debug {
                fmt.Println("x:", x, "y:", y)
            }
            draw(img, c, x,y,*step,*step, *lineWidth)
        }
    }

    // Encode as PNG.
    fmt.Println("Outputting to file: ", strings.TrimSpace(*outputFile))
    f, err := os.OpenFile(*outputFile, os.O_RDWR|os.O_CREATE, 0777)
    if err != nil {
        fmt.Println(err)
    }
    // png.Encode(f, img)
        if err := png.Encode(f, img); err != nil {
        fmt.Println(err)
    }
    f.Close()
}
