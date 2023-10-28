package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	mockdb "simple-gobank/db/mock"
	db "simple-gobank/db/sqlc"
	"simple-gobank/util"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type eqCreateUserParamsMatcher struct {
	arg      db.CreateUserParams
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUserParams)
	if !ok {
		return false
	}

	err := util.CheckPassword(e.password, arg.HashedPassword)

	if err != nil {
		return false
	}

	e.arg.HashedPassword = arg.HashedPassword
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg db.CreateUserParams, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := "/users"
	userRequest := createUserRequest{
		Username: util.RandomString(6),
		Password: util.RandomString(6),
		Fullname: util.RandomString(6),
		Email:    util.RandomEmail(),
	}

	body, _ := json.Marshal(userRequest)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	require.NoError(t, err)
	store.EXPECT().CreateUser(gomock.Any(), EqCreateUserParams(db.CreateUserParams{
		Username: userRequest.Username,
		FullName: userRequest.Fullname,
		Email:    userRequest.Email,
	}, userRequest.Password)).Times(1)
	server.router.ServeHTTP(recorder, request)
}
