package dto

type ListOrganizationsDto struct {
	Page           int    `json:"page,omitempty"`
	Limit          int    `json:"limit,omitempty"`
	FetchAll       bool   `json:"fetchAll,omitempty"`
	WithCustomData bool   `json:"withCustomData,omitempty"`
	TenantId       string `json:"tenantId,omitempty"`
}
