package media

// SearchResultDo 结构体
type SearchResultDo struct {
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 视频信息列表
	ResultList []VideoItemExtDo `json:"result_list,omitempty" xml:"result_list>video_item_ext_do,omitempty"`
	// 调用是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 总视频数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
