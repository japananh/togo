package mock

import (
	"context"
	"errors"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/modules/task/taskmodel"
)

type mockTaskStore struct{}

func NewMockTaskStore() *mockTaskStore {
	return &mockTaskStore{}
}

func (m *mockTaskStore) FindTaskByCondition(
	_ context.Context,
	conditions map[string]interface{},
	_ ...string,
) (*taskmodel.Task, error) {
	if val, ok := conditions["id"]; ok && val.(int) == 1 {
		return &taskmodel.Task{
			Title:       "Task 1",
			Description: "Description 1",
			Status:      taskmodel.Open.String(),
			ParentId:    val.(int),
			CreatedBy:   1,
		}, nil
	}
	return nil, common.ErrRecordNotFound
}

func (m *mockTaskStore) CountUserDailyTask(_ context.Context, createdBy int) (int, error) {
	if createdBy == 1 {
		return 1, nil
	}
	if createdBy == 2 {
		return 2, nil
	}
	return 0, errors.New("invalid task creator")
}

func (m *mockTaskStore) CreateTask(_ context.Context, data *taskmodel.TaskCreate) error {
	data.Id = 2
	return nil
}
