package main

import (
  "fmt"
  "os"
	"os/exec"
	"bufio"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("This tool needs 2 args input and output file path")
    os.Exit(1)
	}

	input := os.Args[1]
	output := os.Args[2]

	filter := "fps=12,split[a][b];[a]palettegen[p];[b][p]paletteuse=dither=none"

	args := []string{
		"-loglevel", "error",
		"-i", input,
		"-filter_complex", filter,
		"-y", output,
	}

	cwd, _ := os.Getwd()

  cmd := exec.Command("ffmpeg", args...)

	cmd.Dir = cwd

	stderr, _ := cmd.StderrPipe()

	cmd.Start()

	stderr_scanner := bufio.NewScanner(stderr)
	for stderr_scanner.Scan() {
		fmt.Println(stderr_scanner.Text())
	}

	cmd.Wait()
}
