package dto


type CheckPermissionsRespDto struct{
    NamespaceCode  string `json:"namespaceCode"`
    Action  string `json:"action"`
    Resource  string `json:"resource"`
    Enabled  bool `json:"enabled"`
}

