package dto


type ListGroupsDto struct{
    Keywords string `json:"keywords,omitempty"`
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

