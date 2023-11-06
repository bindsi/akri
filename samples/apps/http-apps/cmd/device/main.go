package main

import (
  "flag"
  "fmt"
  "log"
  "math/rand"
  "net"
  "net/http"
  "time"
  "strings"
  "html"
)

const (
  addr = ":8080"
)

// RepeatableFlag is an alias to use repeated flags with flag
type RepeatableFlag []string

// String is a method required by flag.Value interface
func (e *RepeatableFlag) String() string {
  result := strings.Join(*e, "\n")
  return result
}

// Set is a method required by flag.Value interface
func (e *RepeatableFlag) Set(value string) error {
  *e = append(*e, value)
  return nil
}
var _ flag.Value = (*RepeatableFlag)(nil)
var paths RepeatableFlag
var devices RepeatableFlag

func main() {
  flag.Var(&paths, "path", "Repeat this flag to add paths for the device")
  flag.Var(&devices, "device", "Repeat this flag to add devices to the discovery service")
  flag.Parse()

  // At a minimum, respond on `/`
  if len(paths) == 0 {
    paths = []string{"/"}
  }
  log.Printf("[main] Paths: %d", len(paths))

  seed := rand.NewSource(time.Now().UnixNano())
  entr := rand.New(seed)

  handler := http.NewServeMux()

  // Create handler for the discovery endpoint
  handler.HandleFunc("/discovery", func(w http.ResponseWriter, r *http.Request) {
    log.Printf("[discovery] Handler entered")
    fmt.Fprintf(w, "%s\n", html.EscapeString(devices.String()))
  })
  // Create handler for each endpoint
  for _, path := range paths {
    log.Printf("[main] Creating handler: %s", path)
    handler.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
      log.Printf("[device] Handler entered: %s", path)
      fmt.Fprint(w, entr.Float64())
    })
  }

  s := &http.Server{
    Addr:    addr,
    Handler: handler,
  }
  listen, err := net.Listen("tcp", addr)
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("[main] Starting Device: [%s]", addr)
  log.Fatal(s.Serve(listen))
}