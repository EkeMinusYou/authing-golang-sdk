package dto


type GetOrdersDto struct{
    Page int `json:"page,omitempty"`
    Limit int `json:"limit,omitempty"`
}

