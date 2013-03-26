package main

import (
  "archive/zip"
  "io/ioutil"
  "log"
  "fmt"
  "os"
  "path/filepath"
)


func zip_dir(dir string, zip_name string) error{

  zip_file, err := os.Create(zip_name) 
  if err != nil {
    return err
  }
  w := zip.NewWriter(zip_file)

  addFile := func (path string, info os.FileInfo, err error) error {

    if err != nil {
      return err
    }

    f, err := w.Create(path); 
    if err != nil {
      return err
    }
    if info.IsDir() {
      return nil
    }

    content, err := ioutil.ReadFile(path);
    if err != nil {
      return err
    }

    _, err = f.Write(content)
    if err != nil {
      return err
    }


    return nil
  }

  if err = filepath.Walk(dir, addFile); err != nil {
    return err
  }
  err = w.Close()
  return err
}



func main() {

  if len(os.Args) < 3 {
    fmt.Println("Please specify directory name and zip anme")
    os.Exit(-1)
  }

  if err := zip_dir(os.Args[1], os.Args[2]);err != nil {
    log.Fatal(err)
    os.Exit(-1)
  }
  os.Exit(0)
}

