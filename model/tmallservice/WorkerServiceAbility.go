package tmallservice

// WorkerServiceAbility 结构体
type WorkerServiceAbility struct {
	// 工人可服务区域
	AreaCodeList []int64 `json:"area_code_list,omitempty" xml:"area_code_list>int64,omitempty"`
}
