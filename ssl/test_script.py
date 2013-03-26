#!/usr/bin/env python
# Usage: python check_agent.py IP 

import socket
import ssl
import sys
import time

s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

def create_socket(host, port):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    while True:
       try:
           s.connect( (host, int(port)) )
       except:
           s.shutdown(socket.SHUT_RDWR)
           s.close()
           time.sleep(1)
       else:
           break
    
    return ssl.wrap_socket(s, keyfile="sample.pem",
                              certfile="sample.pem")


def normal_command(host, port, code):
    ssl_sock = create_socket(host, int(port))
    msg = "33"
    ssl_sock.write(str(len(msg)))
    ssl_sock.write(chr(255)) # end of length
    ssl_sock.write(msg) #id content
    print ssl_sock.read()
    s.close()
  
if __name__ == '__main__':
    try:

        import sys
        if len(sys.argv) <= 1:
            print "IP doesn't specifiy"
            sys.exit(-1)
        ip = sys.argv[1]
        normal_command(ip, 44443, 0)
    except Exception, e:
        sys.stderr.write("Exception: %s" % str(e))
        sys.exit(-1)
    sys.exit(0)


