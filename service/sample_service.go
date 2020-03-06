package service

import (
	"gomock-test/model"
	"gomock-test/repository"
)

type SampleService struct {
	isr repository.ISampleRepository
}

func (ss SampleService) GetSample() model.Sample {
	return ss.isr.GetSample()
}