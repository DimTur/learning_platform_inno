package ssomodels

type CreateLearningGroup struct {
	Name        string   `json:"name" validate:"required,min=3,max=100"`
	CreatedBy   string   `json:"created_by" validate:"required"`
	ModifiedBy  string   `json:"modified_by" validate:"required"`
	GroupAdmins []string `json:"group_admins" validate:"required"`
	Learners    []string `json:"learners" validate:"required"`
}

type CreateLearningGroupResp struct {
	Success bool `json:"success"`
}

type GetLgByID struct {
	UserID string `json:"user_id" validate:"required"`
	LgId   string `json:"learning_group_id" validate:"required"`
}

type GetLgByIDResp struct {
	Id          string         `json:"id"`
	Name        string         `json:"name"`
	CreatedBy   string         `json:"created_by"`
	ModifiedBy  string         `json:"modified_by"`
	Learners    []*Learner     `json:"learners"`
	GroupAdmins []*GroupAdmins `json:"group_admins"`
}

type Learner struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type GroupAdmins struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UpdateLearningGroup struct {
	UserID      string   `json:"user_id" validate:"required"`
	LgId        string   `json:"learning_group_id" validate:"required"`
	Name        string   `json:"name,omitempty"`
	ModifiedBy  string   `json:"modified_by,omitempty"`
	GroupAdmins []string `json:"group_admins,omitempty"`
	Learners    []string `json:"learners,omitempty"`
}

type UpdateLearningGroupResp struct {
	Success bool `json:"success"`
}

type DelLgByID struct {
	UserID string `json:"user_id" validate:"required"`
	LgID   string `json:"learning_group_id" validate:"required"`
}

type DelLgByIDResp struct {
	Success bool `json:"success"`
}

type GetLGroups struct {
	UserID string `json:"user_id" validate:"required"`
}

type GetLGroupsResp struct {
	LearningGroups []*LearningGroup `json:"learning_groups"`
}

type LearningGroup struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

type IsGroupAdmin struct {
	UserID string `json:"user_id" validate:"required"`
	LgID   string `json:"learning_group_id" validate:"required"`
}

type IsGroupAdminResp struct {
	IsGroupAdmin bool `json:"is_group_admin"`
}

type UserIsGroupAdminIn struct {
	UserID string `json:"user_id" validate:"required"`
}

type UserIsLearnerIn struct {
	UserID string `json:"user_id" validate:"required"`
}
