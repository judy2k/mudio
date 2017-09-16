package server

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/satori/go.uuid"
)

type FileMaker struct {
	Output string
	Binary string
	Args   []string
}

// Worker processes individual commands.
func Worker(i int, input <-chan FileMaker) {
	for {
		fm := <-input
		cmd := exec.Command(fm.Binary, fm.Args...)
		cmd.Stderr = os.Stdout
		res := cmd.Run()
		fmt.Println(i, "stopped", res)
	}
}

// Run farms out processes to the workers.
func Run(port int) {
	jobs := make(chan FileMaker, 10)
	for i := 0; i < 5; i++ {
		go Worker(i, jobs)
	}

	for i := 0; i < 20; i++ {
		u1 := uuid.NewV4()
		fn := fmt.Sprintf("/tmp/out-%s.mp3", u1)
		jobs <- FileMaker{
			Output: fn,
			Binary: "sox",
			Args: []string{
				"research/Sample.mp3",
				fn,
				"norm", "reverse", "reverb", "reverse"},
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}
