package wdk

// EbeecakeO2OCallbackContent 结构体
type EbeecakeO2OCallbackContent struct {
	// 状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 作业内容单号
	WorkUnitContentId string `json:"work_unit_content_id,omitempty" xml:"work_unit_content_id,omitempty"`
}
