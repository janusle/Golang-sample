#!/usr/bin/env ruby

require "rubygems"
require "json"

src = { :name => "Bender", :age => 23 }
jsn = JSON.dump(src) # obj -> json

dst = JSON.parse(jsn) # json -> obj
puts dst["name"]
puts dst["age"]



