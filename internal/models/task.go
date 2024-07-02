package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"goEdu/internal/enum/taskStatus"
)

type Task struct {
	UUID               pgtype.UUID           `json:"uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	UserUUID           pgtype.UUID           `json:"user_uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Status             taskStatus.TaskStatus `json:"status" example:"created"`
	Message            string                `json:"message" example:"message"`
	UserUploadLink     string                `json:"user_upload_link" example:"http"`
	UserDownloadLink   string                `json:"user_download_link" example:"http"`
	ServerUploadLink   string                `json:"server_upload_link" example:"http"`
	ServerDownloadLink string                `json:"server_download_link" example:"http"`
	UpdatedAt          pgtype.Timestamptz    `json:"updated_at" example:"2020-01-01T00:00:00Z"`
	CreatedAt          pgtype.Timestamptz    `json:"created_at" example:"2020-01-01T00:00:00Z"`
}
