package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// init logging
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msgf("config load [OK]")

	// init tracing span

}
