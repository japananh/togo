package mock

import (
	"context"
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component/tokenprovider"
	"github.com/japananh/togo/modules/user/usermodel"
	"time"
)

type mockUserStore struct{}

func NewMockUserStore() *mockUserStore {
	return &mockUserStore{}
}

func (mockUserStore) FindUser(_ context.Context, conditions map[string]interface{}, _ ...string) (*usermodel.User, error) {
	if val, ok := conditions["email"]; ok && val.(string) == "user@gmail.com" {
		return &usermodel.User{Email: val.(string), Password: "user@123", DailyTaskLimit: 5, Status: 1, Salt: ""}, nil
	}
	return nil, common.ErrRecordNotFound
}

func (mockUserStore) CreateUser(_ context.Context, data *usermodel.UserCreate) error {
	data.Id = 2
	return nil
}

type mockProvider struct{}

func NewMockProvider() *mockProvider {
	return &mockProvider{}
}

func (mockProvider) Generate(_ tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	return &tokenprovider.Token{
		Token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwYXlsb2FkIjp7InVzZXJfaWQiOjF9LCJleHAiOjE2NTM0MDk1MDksImlhdCI6MTY1MzMyMzEwOX0.SYzR9JXyIc_VeuLXLAnxFWTM3nO6LQfWbyO-vTK3fMo",
		Expiry:  expiry,
		Created: time.Now().UTC(),
	}, nil
}

func (mockProvider) Validate(_ string) (*tokenprovider.TokenPayload, error) {
	return &tokenprovider.TokenPayload{UserId: 1}, nil
}

type mockHash struct{}

func NewMockHash() *mockHash {
	return &mockHash{}
}

func (mockHash) Hash(data string) string {
	return data
}
