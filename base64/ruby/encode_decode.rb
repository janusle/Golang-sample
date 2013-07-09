#!/usr/bin/env ruby

require "base64"

if ARGV.length < 1
  puts "Please specify filename"
  exit 1
end

filename = ARGV[0]

f1 = File.new(filename, 'r')
bin1 = f1.read
st1 = Base64.encode64(bin1)

bin2 = Base64.decode64(st1)
f2 = File.new(filename + ".new", "w")
f2.write(bin2)
f2.close




