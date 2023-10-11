package vaccin

// IsvPovSubscribeInfoResponse 结构体
type IsvPovSubscribeInfoResponse struct {
	// 自有pov预约信息
	InfoDetailList []PovSubscribeDetailModel `json:"info_detail_list,omitempty" xml:"info_detail_list>pov_subscribe_detail_model,omitempty"`
	// 总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}
