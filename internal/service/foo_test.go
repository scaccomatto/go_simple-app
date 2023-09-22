package service

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	mock_database "paru.net/gosimpleapp/internal/test/mocks/database"
)

// We need to mock the dependency(db) if we want to be able to run the tests
func Test_GivenFooWhenDoingSomethingThenNoError(t *testing.T) {
	t.Helper()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dbMock := mock_database.NewMockDatabase(ctrl)

	dbMock.EXPECT().InsertFoo(gomock.Any(), gomock.Any()).Times(1)

	fooSvc := NewFooService(dbMock)
	err := fooSvc.InsertFoo("aa", 1)
	assert.Nil(t, err)
}

func setUpTest() {

}
