package sample

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "github.com/884884/gomock_tutorial/sample/mock_sample"
)

func TestSample1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock.NewMockSample(ctrl)
	mockSample.EXPECT().Method("hoge").Return(1)

	t.Log("result:", mockSample.Method("hoge"))
}