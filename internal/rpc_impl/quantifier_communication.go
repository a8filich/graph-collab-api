package rpc_impl

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/rs/zerolog"

	"github.com/a8filich/graph-collab-api/internal/data"
)

var (
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	host = "[::]"
	port = 50051
)

func SendBig5DataPointToQuantifier(big5DataPoint *data.Big5Data) {
	logger.Debug().Msg("Sending Big5 data point to Quantifier...")

	var opts []grpc.DialOption

	serverAddr := fmt.Sprintf("%s:%d", host, port)

	testCreds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, testCreds)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error dialing to Quantifier: %v", err))
	}
	defer conn.Close()

	client := NewQuantifierCommunicationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	big5Msg := Big5Msg{
		Id: big5DataPoint.UUID.String(),
		Person: &PersonMsg{
			PersonId: big5DataPoint.Person.PersonID,
			Name:     big5DataPoint.Person.Name,
			Age:      uint32(big5DataPoint.Person.Age),
			Sex:      big5DataPoint.Person.Sex,
		},
		Conscientiousness: big5DataPoint.Conscientiousness,
		Openness:          big5DataPoint.Openness,
		Neuroticism:       big5DataPoint.Neuroticism,
		Extraversion:      big5DataPoint.Extraversion,
		Agreeableness:     big5DataPoint.Agreeableness,
	}

	_, err = client.SendBig5Msg(ctx, &big5Msg)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error sending Big5 data point to Quantifier: %v", err))
	}

	logger.Debug().Msg(fmt.Sprintf("Finished sending Big5 data point to Quantifier"))
}

func SendTeamRoleDataPointToQuantifier(teamRoleDataPoint *data.TeamRoleData) {
	logger.Debug().Msg("Sending Team Role data point to Quantifier...")

	var opts []grpc.DialOption

	serverAddr := fmt.Sprintf("%s:%d", host, port)

	testCreds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, testCreds)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error dialing to Quantifier: %v", err))
	}
	defer conn.Close()

	client := NewQuantifierCommunicationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	teamRoleMsg := TeamRoleMsg{
		Id: teamRoleDataPoint.UUID.String(),
		Person: &PersonMsg{
			PersonId: teamRoleDataPoint.Person.PersonID,
			Name:     teamRoleDataPoint.Person.Name,
			Age:      uint32(teamRoleDataPoint.Person.Age),
			Sex:      teamRoleDataPoint.Person.Sex,
		},
		ResourceInvestigator: uint32(teamRoleDataPoint.ResourceInvestigator),
		Teamworker:           uint32(teamRoleDataPoint.Teamworker),
		Coordinator:          uint32(teamRoleDataPoint.Coordinator),
		Plant:                uint32(teamRoleDataPoint.Plant),
		MonitorEvaluator:     uint32(teamRoleDataPoint.MonitorEvaluator),
		Specialist:           uint32(teamRoleDataPoint.Specialist),
		Shaper:               uint32(teamRoleDataPoint.Shaper),
		Implementer:          uint32(teamRoleDataPoint.Implementer),
		CompleterFinisher:    uint32(teamRoleDataPoint.CompleterFinisher),
	}

	_, err = client.SendTeamRoleMsg(ctx, &teamRoleMsg)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error sending Team Role data point to Quantifier: %v", err))
	}

	logger.Debug().Msg(fmt.Sprintf("Finished sending Team Role data point to Quantifier"))
}

func SendExpertiseDataPointToQuantifier(expertiseDataPoint *data.ExpertiseData) {
	logger.Debug().Msg("Sending Expertise data point to Quantifier...")

	var opts []grpc.DialOption

	serverAddr := fmt.Sprintf("%s:%d", host, port)

	testCreds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, testCreds)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error dialing to Quantifier: %v", err))
	}
	defer conn.Close()

	client := NewQuantifierCommunicationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	expertiseMsg := ExpertiseMsg{
		Id: expertiseDataPoint.UUID.String(),
		Person: &PersonMsg{
			PersonId: expertiseDataPoint.Person.PersonID,
			Name:     expertiseDataPoint.Person.Name,
			Age:      uint32(expertiseDataPoint.Person.Age),
			Sex:      expertiseDataPoint.Person.Sex,
		},
		CurrentExpertise: expertiseDataPoint.CurrentExpertise,
		DevelopmentGoals: expertiseDataPoint.DevelopmentGoals,
	}

	_, err = client.SendExpertiseMsg(ctx, &expertiseMsg)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error sending Expertise data point to Quantifier: %v", err))
	}

	logger.Debug().Msg(fmt.Sprintf("Finished sending Expertise data point to Quantifier"))
}

func SendManagerialDataPointToQuantifier(managerialDataPoint *data.ManagerialData) {
	logger.Debug().Msg("Sending Managerial data point to Quantifier...")

	var opts []grpc.DialOption

	serverAddr := fmt.Sprintf("%s:%d", host, port)

	testCreds := grpc.WithTransportCredentials(insecure.NewCredentials())
	opts = append(opts, testCreds)

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error dialing to Quantifier: %v", err))
	}
	defer conn.Close()

	client := NewQuantifierCommunicationClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	managerialMsg := ManagerialMsg{
		Id: managerialDataPoint.UUID.String(),
		Person: &PersonMsg{
			PersonId: managerialDataPoint.Person.PersonID,
			Name:     managerialDataPoint.Person.Name,
			Age:      uint32(managerialDataPoint.Person.Age),
			Sex:      managerialDataPoint.Person.Sex,
		},
		Communication:      uint32(managerialDataPoint.Communication),
		Adaptability:       uint32(managerialDataPoint.Adaptability),
		ConflictResolution: uint32(managerialDataPoint.ConflictResolution),
		DecisionMaking:     uint32(managerialDataPoint.DecisionMaking),
		Leadership:         uint32(managerialDataPoint.Leadership),
		ExperienceYears:    uint32(managerialDataPoint.ExperienceYears),
	}

	_, err = client.SendManagerialMsg(ctx, &managerialMsg)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("Error sending Managerial data point to Quantifier: %v", err))
	}

	logger.Debug().Msg(fmt.Sprintf("Finished sending Managerial data point to Quantifier"))
}
