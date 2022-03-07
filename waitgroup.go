package main
  
import (
  "fmt"
  "io/ioutil"
  "strings"
  "path/filepath"
  "time"
  "sync"
)

var (
  matches []string
  wg = sync.WaitGroup{}
  lock = sync.Mutex{}
)

func main() {
  start := time.Now()
  wg.Add(1)
  fileSearch("/home/cohesity", "waitgroup.go")
  wg.Wait()

  for _, file := range matches {
    fmt.Println("Matched ", file)
  }

  elapsed := time.Since(start)
  fmt.Println("Time took to execute: ", elapsed)
}

func fileSearch(root string, filename string) {
  files, err := ioutil.ReadDir(root)
  if err != nil {
    fmt.Println(err)
    wg.Done()
    return
  }

  for _, file := range files {
    if strings.Contains(file.Name(), "workspace") {
      continue
    }

    if strings.Contains(file.Name(), "software") {
      continue
    }

    fmt.Println(file.Name())
    if strings.Contains(file.Name(), filename) {
      lock.Lock()
      matches = append(matches, filepath.Join(root, file.Name()))
      lock.Unlock()
    }
package main
  
import (
  "fmt"
  "io/ioutil"
  "strings"
  "path/filepath"
  "time"
  "sync"
)

var (
  matches []string
  wg = sync.WaitGroup{}
  lock = sync.Mutex{}
)

func main() {
  start := time.Now()
  wg.Add(1)
  fileSearch("/home/cohesity", "waitgroup.go")
  wg.Wait()

  for _, file := range matches {
    fmt.Println("Matched ", file)
  }

  elapsed := time.Since(start)
  fmt.Println("Time took to execute: ", elapsed)
}

func fileSearch(root string, filename string) {
  files, err := ioutil.ReadDir(root)
  if err != nil {
    fmt.Println(err)
    wg.Done()
    return
  }

  for _, file := range files {
    if strings.Contains(file.Name(), "workspace") {
      continue
    }

    if strings.Contains(file.Name(), "software") {
      continue
    }

    fmt.Println(file.Name())
    if strings.Contains(file.Name(), filename) {
      lock.Lock()
      matches = append(matches, filepath.Join(root, file.Name()))
      lock.Unlock()
    }
    if file.IsDir() {
      wg.Add(1)
      fileSearch(filepath.Join(root, file.Name()), filename)
    }
  }

  defer wg.Done()
}
