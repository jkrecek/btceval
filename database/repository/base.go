package repository

import "github.com/jkrecek/btceval/database"

type Manager struct {
	db *database.DB
}

func NewManager(db *database.DB) *Manager {
	return &Manager{db}
}

func (rm *Manager) Record() RecordRepository {
	return RecordRepository{rm, rm.db}

}
