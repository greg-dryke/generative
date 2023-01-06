#! /bin/bash

# Probably just make this script all lines, one for each type genned for phone/computer
#go build lines.go 

#pixel 3 H: 3120 x W: 1440


#if [ if [ $(ls -A lines_old/pixel3* 2>/dev/null | wc -l) -gt 10 ]
#then
#fi
# Clean up and then store old, this should work! Otherwise i can use above.
ls -1td lines_old/pixel7* | tail -n +10 | xargs rm
mv pixel3.png lines_old/pixel7_$(date "+%Y-%m-%d_%H%M").png 2>/dev/null

## not being pushed?? Change here isn't
# homescreen: ./generative -width 1440 -height 3120 -step 80 -lwidth 13 -output pixel7_homescreen.png -blankZone 0,190,1440,330
./generative -width 1440 -height 3120 -step 80 -lwidth 13 -output pixel7.png -blankZone 0,190,1440,250 -blankZone 140,400,680,970
