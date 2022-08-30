package waybill

// BaseResultDto 结构体
type BaseResultDto struct {
	// 异常信息
	ErrorInfoList []ErrorInfo `json:"error_info_list,omitempty" xml:"error_info_list>error_info,omitempty"`
	// 地址可达结果
	Module *ReachableRecommendResponseDto `json:"module,omitempty" xml:"module,omitempty"`
	// 请求是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
