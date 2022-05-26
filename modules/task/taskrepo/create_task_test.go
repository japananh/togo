package taskrepo_test

import (
	"errors"
	"github.com/japananh/togo/mock"
	"github.com/japananh/togo/modules/task/taskmodel"
	"github.com/japananh/togo/modules/task/taskrepo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateTaskRepo_CreateTask(t *testing.T) {
	var tcs = []struct {
		assignee    int
		createdBy   int
		parentId    int
		title       string
		description string
		expectedErr error
	}{
		{0, 1, 0, "Task 1", "Description 1", nil},
		{0, 2, 0, "Task 2", "Description 1", errors.New("exceed daily task limit")},
		{1, 3, 0, "Task 3", "Description 2", errors.New("invalid task creator")},
		{3, 1, 0, "Task 4", "Description 3", errors.New("invalid assignee")},
		{0, 1, 3, "Task 5", "Description 3", errors.New("invalid parent task")},
	}

	for _, tc := range tcs {
		repo := taskrepo.NewCreateTaskRepo(mock.NewMockTaskStore(), mock.NewMockUserStore())
		err := repo.CreateTask(nil, &taskmodel.TaskCreate{
			Title:       tc.title,
			Description: tc.description,
			AssigneeId:  tc.assignee,
			CreatedBy:   tc.createdBy,
			ParentId:    tc.parentId,
		})
		if tc.expectedErr != nil {
			assert.Error(t, err)
			assert.Equal(t, tc.expectedErr.Error(), err.Error())
		} else {
			assert.Nil(t, err)
		}
	}
}
