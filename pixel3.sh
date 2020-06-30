#! /bin/bash

# Probably just make this script all lines, one for each type genned for phone/computer
#go build lines.go 

#pixel 3 H: 2160 x W: 1080


#if [ if [ $(ls -A lines_old/pixel3* 2>/dev/null | wc -l) -gt 10 ]
#then
#fi
# Clean up and then store old, this should work! Otherwise i can use above.
ls -1td lines_old/* | tail -n +10 | xargs rm
mv pixel3.png lines_old/pixel3_$(date "+%Y-%m-%d_%H%M").png 2>/dev/null

## not being pushed?? Change here isn't
./generative -width 1080 -height 2160 -step 25 -output pixel3.png
