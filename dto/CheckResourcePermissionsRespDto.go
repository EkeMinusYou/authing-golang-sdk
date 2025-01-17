package dto

type CheckResourcePermissionsRespDto struct {
	StatusCode int                            `json:"statusCode"`
	Message    string                         `json:"message"`
	ApiCode    int                            `json:"apiCode,omitempty"`
	Data       CheckResourcePermissionDataDto `json:"data"`
}
