package dto


type ListRoleMembersDto struct{
    Code string `json:"code,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
    WithCustomData bool `json:"withCustomData,omitempty"`
    WithIdentities bool `json:"withIdentities,omitempty"`
    WithDepartmentIds bool `json:"withDepartmentIds,omitempty"`
    Namespace string `json:"namespace,omitempty"`
}

