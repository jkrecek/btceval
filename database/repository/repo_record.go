package repository

import (
	"time"

	"github.com/jkrecek/btceval/database"
	"github.com/jkrecek/btceval/database/entity"

	"github.com/jkrecek/sorm-go"
)

type RecordRepository struct {
	repositoryManager *Manager
	db                *database.DB
}

type scopedRecordRepository struct {
	RecordRepository
	entity *entity.Record
}

func (r RecordRepository) WithEntity(entity *entity.Record) *scopedRecordRepository {
	return &scopedRecordRepository{r, entity}
}

func (r *scopedRecordRepository) IsZero() bool {
	return r.entity == nil || r.entity.ID == 0
}

func (r *scopedRecordRepository) Get() *entity.Record {
	return r.entity
}

func (r *scopedRecordRepository) Save() {
	if r.entity.ID == 0 {
		r.entity.CreatedAt = time.Now()
	}

	sorm.Save(r.db, r.entity)
}
