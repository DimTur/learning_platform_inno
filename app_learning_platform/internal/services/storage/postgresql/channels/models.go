package channels

import (
	"database/sql"
	"time"
)

type Channel struct {
	ID             int64
	Name           string
	Description    string
	CreatedBy      string
	LastModifiedBy string
	CreatedAt      time.Time
	Modified       time.Time
}

type GetChannelByID struct {
	ChannelID int64 `json:"channel_id" validate:"required"`
}

type GetChannels struct {
	LgIDs  []string `json:"learning_group_ids" validate:"required"`
	Limit  int64    `json:"limit,omitempty" validate:"min=1"`
	Offset int64    `json:"offset,omitempty" validate:"min=0"`
}

func (p *GetChannels) SetDefaults() {
	if p.Limit == 0 {
		p.Limit = 10
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
}

type ChannelWithPlans struct {
	ID             int64
	Name           string
	Description    string
	CreatedBy      string
	LastModifiedBy string
	CreatedAt      time.Time
	Modified       time.Time
	Plans          []PlanInChannel
}

type PlanInChannel struct {
	ID             int64
	Name           string
	Description    string
	CreatedBy      string
	LastModifiedBy string
	IsPublished    bool
	Public         bool
	CreatedAt      time.Time
	Modified       time.Time
}

type CreateChannel struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name" validate:"required"`
	Description     string    `json:"description"`
	CreatedBy       string    `json:"created_by" validate:"required"`
	LastModifiedBy  string    `json:"last_modified_by" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	Modified        time.Time `json:"modified"`
	LearningGroupId string    `json:"learning_group_id" validate:"required"`
}

type UpdateChannelRequest struct {
	UserID      string  `json:"user_id" validate:"required"`
	ChannelID   int64   `json:"id" validate:"required"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type DeleteChannelRequest struct {
	ChannelID int64 `json:"id" validate:"required"`
}

type ShareChannelToGroup struct {
	ChannelID int64    `json:"channel_id" validate:"required"`
	LGroupIDs []string `json:"lgroup_ids" validate:"required"`
	CreatedBy string   `json:"created_by" validate:"required"`
}

type IsChannelCreator struct {
	UserID    string `json:"user_id" validate:"required"`
	ChannelID int64  `json:"channel_id" validate:"required"`
}

type DBChannel struct {
	ID             int64     `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	CreatedBy      string    `db:"created_by"`
	LastModifiedBy string    `db:"last_modified_by"`
	CreatedAt      time.Time `db:"created_at"`
	Modified       time.Time `db:"modified"`
}

type DBChannelWithPlans struct {
	ID             int64              `db:"id"`
	Name           string             `db:"name"`
	Description    string             `db:"description"`
	CreatedBy      string             `db:"created_by"`
	LastModifiedBy string             `db:"last_modified_by"`
	CreatedAt      time.Time          `db:"created_at"`
	Modified       time.Time          `db:"modified"`
	Plans          []DBPlanInChannels `db:"plans"`
}

// We can use DBPlan, but for flexibility and control returning fields init new struct
type DBPlanInChannels struct {
	ID             sql.NullInt64  `db:"id"`
	Name           sql.NullString `db:"name"`
	Description    sql.NullString `db:"description"`
	CreatedBy      sql.NullString `db:"created_by"`
	LastModifiedBy sql.NullString `db:"last_modified_by"`
	IsPublished    sql.NullBool   `db:"is_published"`
	Public         sql.NullBool   `db:"public"`
	CreatedAt      sql.NullTime   `db:"created_at"`
	Modified       sql.NullTime   `db:"modified"`
}

type DBShareChannelToGroup struct {
	ChannelID int64     `db:"channel_id"`
	LGroupID  string    `db:"learning_group_id"`
	CreatedBy string    `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
}

type DBChannelCreator struct {
	CreatedBy string `db:"created_by"`
}

type DBLearningGroupID struct {
	LearningGroupID string `db:"learning_group_id"`
}
