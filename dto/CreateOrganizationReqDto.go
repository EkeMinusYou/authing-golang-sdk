package dto

type CreateOrganizationReqDto struct {
	OrganizationName string `json:"organizationName"`
	OrganizationCode string `json:"organizationCode"`
	Description      string `json:"description,omitempty"`
	OpenDepartmentId string `json:"openDepartmentId,omitempty"`
}
