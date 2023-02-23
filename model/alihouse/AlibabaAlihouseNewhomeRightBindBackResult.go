package alihouse

// AlibabaAlihouseNewhomeRightBindBackResult 结构体
type AlibabaAlihouseNewhomeRightBindBackResult struct {
	// 1
	Data []AstoreSceneRespInfoDto `json:"data,omitempty" xml:"data>astore_scene_resp_info_dto,omitempty"`
	// 1
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 1
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 1
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
