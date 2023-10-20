package hotel

// TaobaoxhotelinfolistgetforhelloResultSet 结构体
type TaobaoxhotelinfolistgetforhelloResultSet struct {
	// 每个标准酒店及房型信息集合
	Results []ShotelVo `json:"results,omitempty" xml:"results>shotel_vo,omitempty"`
	// 渠道id
	ChannelId string `json:"channel_id,omitempty" xml:"channel_id,omitempty"`
	// errorCode，error=true时有值
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息，error=true时有值
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// list页h5链接
	H5ListUrl string `json:"h5_list_url,omitempty" xml:"h5_list_url,omitempty"`
	// list页pc链接
	HotelListUrl string `json:"hotel_list_url,omitempty" xml:"hotel_list_url,omitempty"`
	// 请求id，用于链路追踪
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功标志
	SuccessFlag int64 `json:"success_flag,omitempty" xml:"success_flag,omitempty"`
	// 总数据量
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
	// 数据版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 分页查询结果中最后一个id，配合hasNext进行分页查询的性能优化
	LastId int64 `json:"last_id,omitempty" xml:"last_id,omitempty"`
	// 本次请求是否错误
	Error bool `json:"error,omitempty" xml:"error,omitempty"`
	// 当前数据集合是否还有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 是否成功标记
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
