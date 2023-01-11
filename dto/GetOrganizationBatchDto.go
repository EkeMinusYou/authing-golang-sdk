package dto

type GetOrganizationBatchDto struct {
	OrganizationCodeList string `json:"organizationCodeList,omitempty"`
	WithCustomData       bool   `json:"withCustomData,omitempty"`
	TenantId             string `json:"tenantId,omitempty"`
}
