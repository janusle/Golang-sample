package main

import (
  "archive/zip"
  "io"
  "log"
  "fmt"
  "os"
  "bytes"
  "io/ioutil"
)


func UnzipFile(name string) error{
  var (
    err error
    file *os.File
    rc io.ReadCloser
    r *zip.ReadCloser
  )
  if r, err = zip.OpenReader(name); err != nil {
    log.Fatal(err)
    return err
  }
  defer r.Close()

  for _, f := range r.File {
    fmt.Println(f.Name)
    if f.Mode().IsDir() { // if file is a directory
      if err = os.MkdirAll(f.Name, f.Mode().Perm()); err != nil {
        return err
      }
      continue
    }

    if rc, err = f.Open(); err != nil {
      return err
    }

    if file, err = os.Create(f.Name); err != nil {
      return err
    }

    if _, err = io.Copy(file, rc); err != nil {
      return err
    }

    if err = file.Chmod( f.Mode().Perm() ); err != nil {
      return err
    }

    file.Close()
    rc.Close()
  }
  return nil
}


func UnzipMem(arc []byte, size int64) error{
  var (
    err error
    file *os.File
    rc io.ReadCloser
    r *zip.Reader
  )
  ra := bytes.NewReader(arc)

  if r, err = zip.NewReader(ra, size); err != nil {
    log.Fatal(err)
    return err
  }

  for _, f := range r.File {
    fmt.Println(f.Name)
    if f.Mode().IsDir() { // if file is a directory
      if err = os.MkdirAll(f.Name, f.Mode().Perm()); err != nil {
        return err
      }
      continue
    }

    if rc, err = f.Open(); err != nil {
      return err
    }

    if file, err = os.Create(f.Name); err != nil {
      return err
    }

    if _, err = io.Copy(file, rc); err != nil {
      return err
    }

    if err = file.Chmod( f.Mode().Perm() ); err != nil {
      return err
    }

    file.Close()
    rc.Close()
  }
  return nil
}


func main() {
  if len(os.Args) < 3 {
    fmt.Println("Please specify filename")
    os.Exit(-1)
  }

  if os.Args[1] == "-f" {
    // test UnzipFile
    if err := UnzipFile(os.Args[2]); err != nil {
      log.Fatal(err)
    }
  } else {
    // test UnzipMem
    content, err := ioutil.ReadFile(os.Args[2])
    if err != nil {
      log.Fatal(err)
    }

    if err := UnzipMem(content, int64(len(content))); err != nil {
      log.Fatal(err)
    }

  }
  os.Exit(0)  
}

