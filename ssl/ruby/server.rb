#!/usr/bin/env ruby

require "openssl"
require "thread"
require "socket"

PORT = 44443
BUFSIZ = 1024
CERT = "server.pem"

s = TCPServer.new PORT #TCPServer
c = OpenSSL::SSL::SSLContext.new #Context
c.ssl_version = :SSLv23
c.cert = OpenSSL::X509::Certificate.new(File.open(CERT))
c.key = OpenSSL::PKey::RSA.new(File.open(CERT))
ss = OpenSSL::SSL::SSLServer.new(s, c)

loop do
  conn = ss.accept
  Thread.new do
    begin
      while (data = conn.gets)
        puts data
        conn.puts data
      end
    rescue
      $stderr.puts $!
      ss.close
    ensure
      s.close if s
    end
  end
end

