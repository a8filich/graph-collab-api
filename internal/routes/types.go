package routes

import (
	"github.com/google/uuid"
)

type Person struct {
	PersonID string `json:"personID"`
	Name     string `json:"name"`
	Age      uint8  `json:"age"`
	Sex      string `json:"sex"`
}

type Big5Data struct {
	UUID              uuid.UUID `json:"id"`
	Person            Person    `json:"person"`
	Conscientiousness float32   `json:"conscientiousness"`
	Openness          float32   `json:"openness"`
	Neuroticism       float32   `json:"neuroticism"`
	Extraversion      float32   `json:"extraversion"`
	Agreeableness     float32   `json:"agreeableness"`
}

type TeamRoleData struct {
	UUID                 uuid.UUID `json:"id"`
	Person               Person    `json:"person"`
	ResourceInvestigator uint8     `json:"resourceInvestigator"`
	Teamworker           uint8     `json:"teamworker"`
	Coordinator          uint8     `json:"coordinator"`
	Plant                uint8     `json:"plant"`
	MonitorEvaluator     uint8     `json:"monitorEvaluator"`
	Specialist           uint8     `json:"specialist"`
	Shaper               uint8     `json:"shaper"`
	Implementer          uint8     `json:"implementer"`
	CompleterFinisher    uint8     `json:"completerFinisher"`
}

type ExpertiseData struct {
	UUID             uuid.UUID `json:"id"`
	Person           Person    `json:"person"`
	CurrentExpertise []string  `json:"currentExpertise"`
	DevelopmentGoals []string  `json:"developmentGoals"`
}

type ManagerialData struct {
	UUID               uuid.UUID `json:"id"`
	Person             Person    `json:"person"`
	Communication      uint8     `json:"communication"`
	Adaptability       uint8     `json:"adaptability"`
	ConflictResolution uint8     `json:"conflictResolution"`
	DecisionMaking     uint8     `json:"decisionMaking"`
	Leadership         uint8     `json:"leadership"`
	ExperienceYears    uint8     `json:"experienceYears"`
}
