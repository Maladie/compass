package eventapi

import (
	"github.com/kyma-incubator/compass/components/director/internal/model"
)

type inMemoryRepository struct {
	store map[string]*model.EventAPIDefinition
}

func NewRepository() *inMemoryRepository {
	return &inMemoryRepository{store: make(map[string]*model.EventAPIDefinition)}
}

func (inMemoryRepository) ListByApplicationID(applicationID string) ([]*model.EventAPIDefinition, error) {
	//TODO panic("implement me")
	return nil, nil
}

func (inMemoryRepository) CreateMany(items []*model.EventAPIDefinition) error {
	//TODO panic("implement me")
	return nil
}

func (inMemoryRepository) DeleteAllByApplicationID(id string) error {
	//TODO panic("implement me")
	return nil
}
