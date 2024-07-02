package taskStatus

type TaskStatus string

const (
	Created         TaskStatus = "created"
	UserUploading   TaskStatus = "user_uploading"
	Queued          TaskStatus = "queued"
	Processing      TaskStatus = "processing"
	ServerUploading TaskStatus = "server_uploading"
	Completed       TaskStatus = "completed"
	Error           TaskStatus = "error"
)
