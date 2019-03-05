package aptible

import (
	json "encoding/json"
	"fmt"
)

type Deployment struct {
	Id        int64  `json:"id"`
	Status    string `json:"status"`
	GitRef    string `json:"text"`
	UserEmail string `json:"user_email"`
	TriggerID string `json:"trigger_id"`
	CreatedAt string `json:"created_at"`
}

type Service struct {
	Id             int64  `json:"id"`
	Service        string `json:"service"`
	Command        string `json:"command"`
	ContainerCount int64  `json:"container_count"`
	ContainerSize  int64  `json:"container_size"`
}

type App struct {
	ID                  int64      `json:"id"`
	Handle              string     `json:"handle"`
	Status              string     `json:"status"`
	GitRemote           string     `json:"git_remote"`
	LastDeployOperation Deployment `json:"last_deploy_operation"`
	Services            []Service  `json:"services"`
}

type Environment struct {
	Id     string `json:"id"`
	Handle string `json:"handle"`
}

func NewApp(blob []byte) *App {
	var app App
	json.Unmarshal(blob, &app)
	return &app
}

func NewAppList(blob []byte) ([]App, error) {
	var apps []App
	err := json.Unmarshal(blob, &apps)
	return apps, err
}

func (a *App) String() string {
	return fmt.Sprintf("id: %d - handle: %s - status: %s", a.ID, a.Handle, a.Status)
}
