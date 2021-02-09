from sense_hat import SenseHat
import time

s = SenseHat()
s.low_light = True

green = (0, 255, 0)
yellow = (247, 172, 0)
red = (255, 0, 0)
white = (255,255,255)
nothing = (0,0,0)
grey = (69,69,69)

def monster():
    R = red
    O = nothing
    logo = [
    O, O, R, R, R, R, O, O,
    O, R, O, O, O, O, R, O,
    R, O, R, O, O, R, O, R,
    R, O, O, O, O, O, O, R,
    O, R, O, R, R, O, R, O,
    O, O, R, O, O, R, O, O,
    O, O, R, O, O, R, O, O,
    O, O, O, R, R, O, O, O,
    ]
    return logo

def potion():
    G = green
    O = nothing
    logo = [
    O, O, O, G, G, O, O, O,
    O, O, O, G, G, O, O, O,
    O, O, O, G, G, O, O, O,
    O, O, G, G, G, G, O, O,
    O, G, G, G, G, G, G, O,
    O, G, G, G, G, G, G, O,
    O, G, G, G, G, G, G, O,
    O, O, G, G, G, G, O, O,
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
    g = grey
    logo = [
    Y, Y, Y, Y, Y, Y, Y, Y,
    Y, Y, Y, g, Y, Y, Y, Y,
    Y, Y, g, g, Y, Y, Y, Y,
    Y, Y, Y, g, Y, Y, Y, Y,
    Y, Y, Y, g, Y, Y, Y, Y,
    Y, Y, Y, g, Y, Y, Y, Y,
    Y, Y, g, g, g, Y, Y, Y,
    Y, Y, Y, Y, Y, Y, Y, Y,
    ]
    return logo

images = [key,monster,potion,victory]
count = 0

while True:
    s.set_pixels(images[count % len(images)]())
    time.sleep(3)
    count += 1