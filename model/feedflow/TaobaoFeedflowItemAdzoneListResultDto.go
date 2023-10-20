package feedflow

// TaobaofeedflowitemadzonelistResultDto 结构体
type TaobaofeedflowitemadzonelistResultDto struct {
	// 广告位列表
	AdzoneList []AdzoneDto `json:"adzone_list,omitempty" xml:"adzone_list>adzone_dto,omitempty"`
	// 成功
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用是否成功,true-成功，false-失败
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
