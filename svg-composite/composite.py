#!/usr/bin/env python

from PIL import Image

hair1 = Image.open("hair1.png")
hair2 = Image.open("hair2.png")

i = 0
for hair in [hair1, hair2]:
    i += 1
    bg = Image.open("face.png")
    bg.paste(hair, (0, 0), hair)
    bg.save("comp{}.png".format(i))

print "Done"
