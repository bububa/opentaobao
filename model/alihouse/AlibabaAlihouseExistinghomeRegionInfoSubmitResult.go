package alihouse

// AlibabaAlihouseExistinghomeRegionInfoSubmitResult 结构体
type AlibabaAlihouseExistinghomeRegionInfoSubmitResult struct {
	// 成功/失败结果返回
	Data []RegionExpertResultDto `json:"data,omitempty" xml:"data>region_expert_result_dto,omitempty"`
	// 处理描述信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 成功/失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
