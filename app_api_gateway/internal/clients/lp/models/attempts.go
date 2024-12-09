package lpmodels

type TryLesson struct {
	UserID    string `json:"user_id" validate:"required"`
	LessonID  int64  `json:"lesson_id" validate:"required"`
	PlanID    int64  `json:"plan_id" validate:"required"`
	ChannelID int64  `json:"channel_id" validate:"required"`
}

type QuestionPageAttempt struct {
	ID              int64  `json:"id" redis:"id"`
	PageID          int64  `json:"page_id" redis:"page_id"`
	LessonAttemptID int64  `json:"lesson_attempt_id" redis:"lesson_attempt_id"`
	IsCorrect       bool   `json:"is_correct" redis:"is_correct"`
	UserAnswer      string `json:"user_answer,omitempty" redis:"user_answer"`
}

type TryLessonResp struct {
	QuestionPageAttempts []QuestionPageAttempt `json:"question_page_attempts"`
}

type UpdatePageAttempt struct {
	UserID          string `json:"user_id" validate:"required"`
	LessonAttemptID int64  `json:"lesson_attempt_id" validate:"required"`
	PageID          int64  `json:"page_id" validate:"required"`
	QPAttemptID     int64  `json:"question_page_attempt_id" validate:"required"`
	UserAnswer      string `json:"user_answer,omitempty"`
}

type UpdatePageAttemptResp struct {
	Success bool `json:"success"`
}

type CompleteLesson struct {
	UserID          string `json:"user_id" validate:"required"`
	LessonAttemptID int64  `json:"lesson_attempt_id" validate:"required"`
}

type CompleteLessonResp struct {
	ID              int64 `json:"id"`
	IsSuccessful    bool  `json:"is_successful"`
	PercentageScore int64 `json:"percentage_score"`
}

type GetLessonAttempts struct {
	UserID   string `json:"user_id" validate:"required"`
	LessonID int64  `json:"lesson_id,omitempty"`
	Limit    int64  `json:"limit,omitempty" validate:"min=1"`
	Offset   int64  `json:"offset,omitempty" validate:"min=0"`
}

type LessonAttempt struct {
	ID              int64  `json:"id"`
	UserID          string `json:"user_id"`
	LessonID        int64  `json:"lesson_id"`
	PlanID          int64  `json:"plan_id"`
	ChannelID       int64  `json:"channel_id"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	IsComplete      bool   `json:"is_complete"`
	IsSuccessful    bool   `json:"is_successful"`
	PercentageScore int64  `json:"percentage_score"`
}

type GetLessonAttemptsResp struct {
	LessonAttempts []LessonAttempt `json:"lesson_attempts"`
}

type LessonAttemptPermissions struct {
	UserID          string `json:"user_id"`
	LessonAttemptID int64  `json:"lesson_attempt_id"`
}
