package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

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

// TODO: implement input data processing
func InputDataBig5Handler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Big Five data")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error().Msg("Error reading request body for Big5 input data")
		http.Error(writer, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var inputData Big5Data
	if err = json.Unmarshal(body, &inputData); err != nil {
		logger.Error().Msg("Error unmarshalling request body for Big5 input data")
		http.Error(writer, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	logger.Info().Msg(fmt.Sprintf("Received Big Five data: %+v", inputData))

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Big5 data")
}

// TODO: implement input data processing
func InputDataTeamRoleHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Team Role data")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error().Msg("Error reading request body for Team Role input data")
		http.Error(writer, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var inputData TeamRoleData
	if err = json.Unmarshal(body, &inputData); err != nil {
		logger.Error().Msg("Error unmarshalling request body for Team Role input data")
		http.Error(writer, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	logger.Info().Msg(fmt.Sprintf("Received Team Role data: %+v", inputData))

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Team Role data")
}

// TODO: implement input data processing
func InputDataExpertiseHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Expertise data")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error().Msg("Error reading request body for Expertise input data")
		http.Error(writer, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var inputData ExpertiseData
	if err = json.Unmarshal(body, &inputData); err != nil {
		logger.Error().Msg("Error unmarshalling request body for Expertise input data")
		http.Error(writer, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	logger.Info().Msg(fmt.Sprintf("Received Expertise data: %+v", inputData))

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Expertise data")
}

// TODO: implement input data processing
func InputDataManagerialHandler(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		logger.Error().Msg(fmt.Sprintf("Incoming %s request, method not allowed", req.Method))
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	logger.Debug().Msg("Incoming request with Managerial data")

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error().Msg("Error reading request body for Managerial input data")
		http.Error(writer, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var inputData ManagerialData
	if err = json.Unmarshal(body, &inputData); err != nil {
		logger.Error().Msg("Error unmarshalling request body for Managerial input data")
		http.Error(writer, "Error unmarshalling request body", http.StatusBadRequest)
		return
	}

	logger.Info().Msg(fmt.Sprintf("Received Managerial data: %+v", inputData))

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintln(writer, "Received Managerial data")
}
