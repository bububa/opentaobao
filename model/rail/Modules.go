package rail

// Modules 结构体
type Modules struct {
	// 坐席code
	SeatCode string `json:"seat_code,omitempty" xml:"seat_code,omitempty"`
	// 坐席详情
	SeatDetail string `json:"seat_detail,omitempty" xml:"seat_detail,omitempty"`
	// 坐席图片
	SeatImage string `json:"seat_image,omitempty" xml:"seat_image,omitempty"`
	// 坐席名称
	SeatName string `json:"seat_name,omitempty" xml:"seat_name,omitempty"`
	// 业务类型，6代表境外火车票
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}
