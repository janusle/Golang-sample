#!/usr/bin/env python

import json

src = { "name": "Bender", "age": 23 }

jsn = json.dumps(src)

dst = json.loads(jsn)

print dst["name"]
print dst["age"]


