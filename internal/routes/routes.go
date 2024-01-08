package routes

import (
	"fmt"
	"net/http"
	"os"
	"time"
	//"io/ioutil"

	"github.com/rs/zerolog"
)

const apiPrefixV1 string = "/api/v1"

var logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Logger()

func Router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/healthz", HealthzHandler)

	mux.HandleFunc(fmt.Sprintf("%s%s", apiPrefixV1, "/input-data/big5"), InputDataBig5Handler)
	mux.HandleFunc(fmt.Sprintf("%s%s", apiPrefixV1, "/input-data/team-role"), InputDataTeamRoleHandler)
	mux.HandleFunc(fmt.Sprintf("%s%s", apiPrefixV1, "/input-data/expertise"), InputDataExpertiseHandler)
	mux.HandleFunc(fmt.Sprintf("%s%s", apiPrefixV1, "/input-data/managerial"), InputDataManagerialHandler)

	return mux
}

func RootHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Info().Msg("Incoming request to the root")

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Hello from Graph Collab")
}

func HealthzHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming health check request")

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Healthy")
}

func InputDataBig5Handler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Big Five data")

	// TODO: Implement request body processing.
	/*
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logger.Error().Msg("Error reading request body for Big5 input data")
			http.Error(writer, "Error reading request body", http.StatusInternalServerError)
			return
		}
	*/

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Big5 data")
}

func InputDataTeamRoleHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Team Role data")

	// TODO: Implement request body processing.
	/*
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logger.Error().Msg("Error reading request body for Team Role input data")
			http.Error(writer, "Error reading request body", http.StatusInternalServerError)
			return
		}
	*/

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Team Role data")
}

func InputDataExpertiseHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Expertise data")

	// TODO: Implement request body processing.
	/*
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logger.Error().Msg("Error reading request body for Expertise input data")
			http.Error(writer, "Error reading request body", http.StatusInternalServerError)
			return
		}
	*/

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Expertise data")
}

func InputDataManagerialHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Managerial data")

	// TODO: Implement request body processing.
	/*
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logger.Error().Msg("Error reading request body for Managerial input data")
			http.Error(writer, "Error reading request body", http.StatusInternalServerError)
			return
		}
	*/

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Managerial data")
}
