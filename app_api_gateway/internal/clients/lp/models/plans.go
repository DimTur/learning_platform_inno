package lpmodels

type CreatePlan struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description,omitempty"`
	CreatedBy       string `json:"created_by" validate:"required"`
	ChannelID       int64  `json:"channel_id" validate:"required"`
	LearningGroupId string `json:"learning_group_id" validate:"required"`
}

type CreatePlanResponse struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success"`
}

type GetPlan struct {
	UserID    string `json:"user_id" validate:"required"`
	PlanID    int64  `json:"plan_id" validate:"required"`
	ChannelID int64  `json:"channel_id" validate:"required"`
}

type GetPlanResponse struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	CreatedBy      string `json:"created_by"`
	LastModifiedBy string `json:"last_modified_by"`
	IsPublished    bool   `json:"is_published"`
	Public         bool   `json:"public"`
	CreatedAt      string `json:"created_at"`
	Modified       string `json:"modified"`
}

type GetPlans struct {
	UserID    string `json:"user_id" validate:"required"`
	ChannelID int64  `json:"channel_id" validate:"required"`
	Limit     int64  `json:"limit,omitempty" validate:"min=1"`
	Offset    int64  `json:"offset,omitempty" validate:"min=0"`
}

type UpdatePlan struct {
	ChannelID      int64   `json:"channel_id" validate:"required"`
	PlanID         int64   `json:"plan_id" validate:"required"`
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	LastModifiedBy string  `json:"last_modified_by" validate:"required"`
	IsPublished    *bool   `json:"is_published,omitempty"`
	Public         *bool   `json:"public,omitempty"`
}

type UpdatePlanResponse struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success"`
}

type DelPlan struct {
	UserID    string `json:"user_id" validate:"required"`
	ChannelID int64  `json:"channel_id" validate:"required"`
	PlanID    int64  `json:"plan_id" validate:"required"`
}

type DelPlanResponse struct {
	Success bool `json:"success"`
}

type SharePlan struct {
	UserID    string   `json:"user_id" validate:"required"`
	ChannelID int64    `json:"channel_id" validate:"required"`
	PlanID    int64    `json:"plan_id" validate:"required"`
	UsersIDs  []string `json:"user_ids" validate:"required"`
}

type SharingPlanResp struct {
	Success bool `json:"success"`
}

type IsPlanShareWith struct {
	IsShare bool `json:"is_share"`
}
