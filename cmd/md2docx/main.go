 package main

 import (
     "flag"
     "fmt"
     "os"
     "os"
     "os/exec"
     "path/filepath"
 )

 func main() {
     template := flag.String("template", "", "Path to reference DOCX template file")
     verbose := flag.Bool("verbose", false, "Enable verbose output")
     flag.Usage = func() {
         fmt.Fprintf(os.Stderr, "Usage: %s [options] <input.md>\n", os.Args[0])
         flag.PrintDefaults()
     }
     flag.Parse()

     if flag.NArg() != 1 {
         flag.Usage()
         os.Exit(1)
     }
     input := flag.Arg(0)

     if _, err := os.Stat(input); os.IsNotExist(err) {
         fmt.Fprintf(os.Stderr, "Error: input file does not exist: %s\n", input)
         os.Exit(1)
     } else if err != nil {
         fmt.Fprintf(os.Stderr, "Error: unable to access input file: %v\n", err)
         os.Exit(1)
     }

     ext := filepath.Ext(input)
     base := input[:len(input)-len(ext)]
     output := base + ".docx"

     args := []string{}
     if *template != "" {
         args = append(args, "--reference-doc", *template)
     }
     args = append(args, "--standalone", "-o", output, input)

     if *verbose {
         fmt.Fprintf(os.Stderr, "Running command: pandoc %v\n", args)
     }
     cmd := exec.Command("pandoc", args...)
     cmd.Stdout = os.Stdout
     cmd.Stderr = os.Stderr
     if err := cmd.Run(); err != nil {
         fmt.Fprintf(os.Stderr, "Error: pandoc conversion failed: %v\n", err)
         os.Exit(1)
     }
     if *verbose {
         fmt.Fprintf(os.Stderr, "Generated %s\n", output)
     }
 }
