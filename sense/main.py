from sense_hat import SenseHat
import time

s = SenseHat()
s.low_light = True

green = (0, 255, 0)
yellow = (255, 255, 0)
red = (255, 0, 0)
white = (255,255,255)
nothing = (0,0,0)

def monster():
    R = red
    O = nothing
    logo = [
    O, O, R, R, R, R, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    ]
    return logo

def potion():
    G = green
    O = nothing
    logo = [
    O, O, O, G, G, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    ]
    return logo

def key():
    W = white
    O = nothing
    logo = [
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O,
    O, W, W, W, O, O, O, O, 
    O, W, O, W, W, W, W, W,
    O, W, W, W, O, O, W, W,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    O, O, O, O, O, O, O, O,
    ]
    return logo

def victory():
    Y = yellow
    O = nothing
    logo = [
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O, 
    O, O, O, O, O, O, O, O,
    ]
    return logo

images = [key,monster,potion,victory]
count = 0

while True: 
    s.set_pixels(images[count % len(images)]())
    time.sleep(3)
    count += 1
