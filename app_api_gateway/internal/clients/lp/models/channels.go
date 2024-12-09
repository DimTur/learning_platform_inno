package lpmodels

type CreateChannel struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description,omitempty"`
	CreatedBy       string `json:"created_by" validate:"required"`
	LearningGroupId string `json:"learning_group_id" validate:"required"`
}

type CreateChannelResponse struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success"`
}

type GetChannel struct {
	UserID    string `json:"user_id" validate:"required"`
	ChannelID int64  `json:"channel_id" validate:"required"`
}

type GetChannelFull struct {
	UserID           string   `json:"user_id" validate:"required"`
	ChannelID        int64    `json:"channel_id" validate:"required"`
	LearningGroupIds []string `json:"learning_group_ids" validate:"required"`
}

type GetChannelResponse struct {
	Id             int64   `json:"id"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	CreatedBy      string  `json:"created_by"`
	LastModifiedBy string  `json:"last_modified_by"`
	CreatedAt      string  `json:"created_at"`
	Modified       string  `json:"modified"`
	Plans          []*Plan `json:"plans"`
}

type Plan struct {
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

type GetChannels struct {
	UserID string `json:"user_id" validate:"required"`
	Limit  int64  `json:"limit,omitempty" validate:"min=1"`
	Offset int64  `json:"offset,omitempty" validate:"min=0"`
}

type GetChannelsFull struct {
	UserID           string   `json:"user_id" validate:"required"`
	LearningGroupIds []string `json:"learning_group_ids" validate:"required"`
	Limit            int64    `json:"limit,omitempty" validate:"min=1"`
	Offset           int64    `json:"offset,omitempty" validate:"min=0"`
}

type Channel struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	CreatedBy      string `json:"created_by"`
	LastModifiedBy string `json:"last_modified_by"`
	CreatedAt      string `json:"created_at"`
	Modified       string `json:"modified"`
}

type UpdateChannel struct {
	UserID      string  `json:"user_id" validate:"required"`
	ChannelID   int64   `json:"id" validate:"required"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type UpdateChannelFull struct {
	UserID       string   `json:"user_id" validate:"required"`
	AdminInLgIds []string `json:"admin_in_lg_ch_ids" validate:"required"`
	ChannelID    int64    `json:"id" validate:"required"`
	Name         *string  `json:"name,omitempty"`
	Description  *string  `json:"description,omitempty"`
}

type UpdateChannelResponse struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success"`
}

type DelChByID struct {
	UserID    string `json:"user_id" validate:"required"`
	ChannelID int64  `json:"id" validate:"required"`
}

type DelChByIDFull struct {
	UserID       string   `json:"user_id" validate:"required"`
	ChannelID    int64    `json:"id" validate:"required"`
	AdminInLgIds []string `json:"admin_in_lg_ch_ids" validate:"required"`
}

type DelChByIDResp struct {
	Success bool `json:"success"`
}

type SharingChannel struct {
	UserID    string   `json:"user_id" validate:"required"`
	ChannelID int64    `json:"channel_id" validate:"required"`
	LGroupIDs []string `json:"lgroup_ids" validate:"required"`
}

type SharingChannelFull struct {
	UserID                       string   `json:"user_id" validate:"required"`
	ChannelID                    int64    `json:"channel_id" validate:"required"`
	LGroupIDs                    []string `json:"lgroup_ids" validate:"required"`
	UserAdminInLearningGroupsIDs []string `json:"user_admin_in_learning_group_ids" validate:"required"`
}

type SharingChannelResp struct {
	Success bool `json:"success"`
}

type IsChannelCreator struct {
	UserID    string `json:"user_id" validate:"required"`
	ChannelID int64  `json:"channel_id" validate:"required"`
}

type IsChannelCreatorResp struct {
	IsCreator bool `json:"is_creator"`
}

type LerningGroupsShareWithChannel struct {
	ChannelID int64 `json:"channel_id" validate:"required"`
}

type LerningGroupsShareWithChannelResp struct {
	LearningGroupIDs []string `json:"learning_group_ids"`
}
