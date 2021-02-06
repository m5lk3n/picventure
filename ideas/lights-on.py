#!/usr/bin/python3

import sys, getopt, os
from sense_hat import SenseHat
import time
from PIL import Image

# from https://magpi.raspberrypi.org/articles/pixel-art-on-sense-hat
def show(color_image):

   # read 8x8 pixel image
   image_file = os.path.join(os.sep,"/home","michael","bin","conditions-to-school",color_image)
   img = Image.open(image_file)

   # generate RGB values for image pixels
   rgb_img = img.convert('RGB')
   image_pixels = list(rgb_img.getdata())

   # extract 8x8=64 pixels
   pixel_width = 1
   image_width = pixel_width*8
   sense_pixels = []
   start_pixel = 0
   while start_pixel < (image_width*64):
      sense_pixels.extend(image_pixels[start_pixel:(start_pixel+image_width):pixel_width])
      start_pixel += (image_width*pixel_width)

   # display image
   sense = SenseHat()
   sense.clear()
   #sense.set_rotation(r=180)
   sense.set_pixels(sense_pixels)
   print ('Resulting indicator image is:', color_image)
   #print(sense_pixels)

def main(argv):
   forecolor = ''
   backcolor = ''
   try:
      opts, args = getopt.getopt(argv,"hf:b:",["forecolor=","backcolor="])
   except getopt.GetoptError:
      print ('lights-on.py -f <color> -b <color>')
      sys.exit(2)
   for opt, arg in opts:
      if opt == '-h':
         print ('lights-on.py -f <color> -b <color>')
         sys.exit()
      elif opt in ("-f", "--forecolor"):
         forecolor = arg
      elif opt in ("-b", "--backcolor"):
         backcolor = arg
   color_image = forecolor + "-on-" + backcolor + ".png"
   #print ('forecolor value is:', forecolor)
   #print ('backcolor value is:', backcolor)
   show(color_image) 

if __name__ == "__main__":
   main(sys.argv[1:])
