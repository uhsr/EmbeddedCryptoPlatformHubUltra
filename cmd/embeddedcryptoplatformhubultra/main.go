// cmd/embeddedcryptoplatformhubultra/main.go
package main

import (
"flag"
"log"
"os"

"embeddedcryptoplatformhubultra/internal/embeddedcryptoplatformhubultra"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := embeddedcryptoplatformhubultra.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
