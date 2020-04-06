#Yes i know this is horribly innefficient, once i wrap my head around numpy stuff ill use it to more efficently fill up the textbox, which is where most of the time consumption is coming from

from PIL import Image, ImageFont, ImageDraw
import textwrap
import argparse
import numpy as np


#Handle args
parser = argparse.ArgumentParser()
parser.add_argument("-t", "--text", help="The text to put in the image")
args = parser.parse_args()

text = '"' + args.text + '"'

font = ImageFont.truetype("static/fonts/textile.ttf", 96)

textLayer = Image.new("RGBA", (1920, 1080), 0)
textDraw = ImageDraw.Draw(textLayer)

space = 0
for line in textwrap.wrap(text, 30):
    lineDimensions = textDraw.textsize(line, font)
    drawPos = ((textLayer.width - lineDimensions[0])//2, space)
    textDraw.text(drawPos, line, "white", font)
    space += 100

textbox = textLayer.crop(textLayer.getbbox())
for pixRow in range(textbox.width):
    for pixel in range(textbox.height):
        if textbox.getpixel((pixRow, pixel)) == (0, 0, 0, 0):
            textbox.putpixel((pixRow,pixel), (0, 0, 0, 255))


finalImage = textLayer = Image.new("RGBA", (1920, 1080), "black")

center = ((finalImage.width-textbox.width)//2, (finalImage.height-textbox.height)//2)
finalImage.paste(textbox, center)

finalImage.save("IASIP.png")