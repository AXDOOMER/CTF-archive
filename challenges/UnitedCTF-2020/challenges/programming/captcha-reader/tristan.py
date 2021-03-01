from pwn import *
import base64
import cv2
import numpy as np
import requests
import json

connection = remote("challenges.unitedctf.ca",4003)
for i in range(10):
  print("try " + str(i))
  bytes = connection.recv()
  print("len: " + str(len(bytes)))
  print("payload: " + str(bytes))
  number_bytes = int(bytes.decode("utf-8").rstrip())
  print(text)
  bytes = connection.recvn(number_bytes)
  print(bytes)

  image = base64.b64decode(bytes.decode("utf-8"))
  with open("captcha.png", "wb+") as f:
    f.write(image)
    f.close()

  image = cv2.imread("captcha.png")
  hsv = cv2.cvtColor(image, cv2.COLOR_BGR2HSV)

  lower_red = np.array([6,190,200])
  upper_red = np.array([24,255,255])

  mask = cv2.inRange(hsv, lower_red, upper_red)
  #res = cv2.bitwise_and(image,image, mask= mask)

  cv2.imwrite('red.png',mask)
  #details = pytesseract.image_to_data(mask)
  #print(details)

  payload = {
              'isOverlayRequired': False,
              'apikey': "f3fcac1ce988957",
              'language': "eng",
            }

  with open('red.png', 'rb') as f:
    r = requests.post('https://api.ocr.space/parse/image',
                      files={'red.png': f},
                      data=payload,
                      )
    answer = r.json()["ParsedResults"][0]["ParsedText"]
    connection.sendline(answer.encode())


  print(connection.recvline())

print(connection.recvline())
print(connection.recvline())
print(connection.recvline())
print(connection.recvline())
print(connection.recvline())
