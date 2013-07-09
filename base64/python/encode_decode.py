#!/usr/bin/env python

import base64
import sys

if len(sys.argv) < 2:
    print "Please specify filename"
    sys.exit(1)

filename = sys.argv[1]
f1 = open(filename, 'r')
bin1 = f1.read()
st = base64.encodestring(bin1)

bin2 = base64.decodestring(st)
f2 = open(filename + ".new", "w")
f2.write(bin2)
f1.close()
f2.close()







