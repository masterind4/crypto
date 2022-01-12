#!/bin/env python3

import sys
import os
import hashlib


def readfile(f):
    with open(f, mode='rb') as file:  # b is important -> binary
        fileContent = file.read()
    return fileContent


def usage():
    print("Usage: {} file target_md5".format(sys.argv[0]))


def compute_md5(data):
    return hashlib.md5(data).hexdigest()


def getRandomEndOfFile():
    ba = bytearray(0x00)
    for i in range(0, 255):
        ba.pop()
        ba.append(i)


if __name__ == "__main__":
    if "--help" in sys.argv or "-h" in sys.argv:
        usage()
        sys.exit(0)

    if len(sys.argv) != 3:
        usage()
        sys.exit(1)

    fileName = sys.argv[1]
    targetMD5 = sys.argv[2]

    fileContent = bytearray(readfile(fileName))

    print("Target md5: {}, Current md5: {}".format(
        targetMD5, compute_md5(fileContent)))

    randomBytes = bytearray(os.urandom(1000))
    fileContent += randomBytes

    with open(fileName + "_appended", "wb") as outputFile:
        outputFile.write(fileContent)
