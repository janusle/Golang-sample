package main

import (
  "fmt"
  "crypto/tls"
  "os"
  "net"
  "strconv"
)


func checkError( err error ) {
  if err != nil {
    fmt.Println(err)
    os.Exit(-1)
  }
}


func try_get_datasize( data *[]byte ) (int, int, error){
  var i int
  size := len(*data)
  for i=0; i<size; i++ {
    if (*data)[i] == byte(255) {
      break
    }
  }

  if i != size {
    n, err := strconv.Atoi(string((*data)[0:i]))
    return n, i, err
  }
  return -1, -1, nil
}


func read_data( conn *net.Conn, data *[]byte ) error {

  var (
    err error
    n int
    data_size = -1
    begin = -1
    total = 0
    buf [32]byte
  )

  // read data
  for {
    n, err = (*conn).Read(buf[0:])
    if err != nil {
      return err 
    }
    *data = append(*data, buf[0:n]...)
    total += n
    if data_size == -1 {
      data_size, begin, err = try_get_datasize(data)
      if err != nil {
         return err
      }
    }
    if data_size != -1 && begin != -1 && total-begin-1 == data_size {
      break
    }
  }
  *data = (*data)[begin+1:]
  return nil
}


func write_data( conn *net.Conn, data *[]byte ) error{
  var (
    size = len(*data)
    written = 0
  )
  for {
    n, err := (*conn).Write(*data)
    if err != nil {
      return err 
    }
    written += n
    if written == size {
      break
    }
  }

  return nil
}


func square( conn *net.Conn ) {
    defer (*conn).Close()
    data := make([]byte, 0)
    err := read_data(conn, &data)
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println("Calculate squre of " + string(data)) 
    n, err := strconv.Atoi(string(data))
    if err != nil {
      fmt.Println(err)
      return
    }
    n *= n
    data = []byte(strconv.Itoa(n))
    write_data(conn, &data) 
}


func main() {
  cert, err := tls.LoadX509KeyPair("zagent.pem", "zagent.pem")
  checkError(err)
  config := tls.Config{Certificates: []tls.Certificate{cert}}
  listener, err := tls.Listen("tcp", ":44443", &config)
  checkError(err)
  fmt.Println("Server is started")
  for {
    conn, err := listener.Accept() 
    if err != nil {
      continue
    }
    go square(&conn)
  }
}
