package pkg

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type testTask struct {
	ID     string `json:"id"`
	Status Status `json:"status"`
	Result string `json:"result,omitempty"`
}

func TestCreateTask(t *testing.T) {
	service := NewService()
	handler := NewHandler(service)

	req := httptest.NewRequest(http.MethodPost, "/tasks", nil)
	w := httptest.NewRecorder()

	handler.CreateTask(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var task testTask
	if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
		t.Fatalf("failed to decode task: %v", err)
	}

	if task.ID == "" {
		t.Error("expected task ID to be set")
	}

	if task.Status != StatusPending && task.Status != StatusRunning {
		t.Errorf("expected status pending or running, got %s", task.Status)
	}
}

func TestGetTask(t *testing.T) {
	service := NewService()
	handler := NewHandler(service)

	task := service.CreateTask()

	time.Sleep(100 * time.Millisecond)

	req := httptest.NewRequest(http.MethodGet, "/tasks/"+task.ID, nil)
	w := httptest.NewRecorder()

	handler.GetTask(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var got testTask
	if err := json.NewDecoder(resp.Body).Decode(&got); err != nil {
		t.Fatalf("failed to decode task: %v", err)
	}

	if got.ID != task.ID {
		t.Errorf("expected task ID %s, got %s", task.ID, got.ID)
	}
}
