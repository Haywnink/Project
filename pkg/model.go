package pkg

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
)

type Task struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
	Result string `json:"result,omitempty"`
}
