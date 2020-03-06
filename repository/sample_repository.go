package repository

import "gomock-test/model"

type ISampleRepository interface {
	GetSample() model.Sample
}

type SampleRepository struct {}

func (sr SampleRepository) GetSample(type ) model.Sample {
return model.Sample{Id: 1, Text: "sample"}
}