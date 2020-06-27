// from https://yourbasic.org/golang/create-image/

package main

import (
    "image"
    "image/color"
    "image/png"
    "os"
    "fmt"
    "math/rand"
)

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
    // TODO move to cmdline arg
    // pixel 3 H: 2160 x W: 1080
width := 1080
height := 2160

upLeft := image.Point{0, 0}
lowRight := image.Point{width, height}

img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

// Colors are defined by Red, Green, Blue, Alpha uint8 values.
cyan := color.RGBA{100, 200, 200, 0xff}

/*
     JS Snippet from: https://generativeartistry.com/tutorials/tiled-lines/
    var canvas = document.querySelector('canvas');
    var context = canvas.getContext('2d');
    
    var size = window.innerWidth;
    var step = 20;
    var dpr = window.devicePixelRatio;
    canvas.width = size * dpr;
    canvas.height = size * dpr;
    context.scale(dpr, dpr);
    
    context.lineCap = 'square';
    context.lineWidth = 2;
    
    function draw(x, y, width, height) {
      var leftToRight = Math.random() >= 0.5;
    
      if(leftToRight) {
        context.moveTo(x, y);
        context.lineTo(x + width, y + height);    
      } else {
        context.moveTo(x + width, y);
        context.lineTo(x, y + height);
      }
    
      context.stroke();
    }
    
    for(var x = 0; x < size; x += step) {
      for(var y = 0; y < size; y+= step) {
        draw(x, y, step, step);    
      }
    }
*/

// Set color for each pixel.
xstep := 10
ystep := 10
for x := 0; x < width; x+=xstep {
    for y := 0; y < height; y+=ystep {
        fmt.Println("x:", x, "y:", y)
        draw(img, cyan, x,y,ystep,xstep)
    }
}

// Encode as PNG.
f, _ := os.Create("image.png")
png.Encode(f, img)
}
