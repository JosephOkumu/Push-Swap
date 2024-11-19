package lib
import (
	"flag"
)

// Determines if the --generate-random flag is set
func EnableRandFlag() bool {
	var generateRandom = flag.Bool("generate-random", false, "Set this flag to enable random number generation")
	flag.Parse()
	return *generateRandom
}
