package gorm

import "testing"

func TestCreateTable(t *testing.T) {
	CreateTable()
}

func TestInsertData(t *testing.T) {
	InsertData()
}

func TestInsertData2(t *testing.T) {
	InsertData2()
}

func TestConflictInsert(t *testing.T) {
	ConflictInsert()
}

func TestConflictInsertComment(t *testing.T) {
	ConflictInsertComment()
}

func TestQueryData(t *testing.T) {
	QueryData()
}
func TestCreatPost(t *testing.T) {
	CreatPost()
}

func TestDeleteComment(t *testing.T) {
	DeleteComment()
}
