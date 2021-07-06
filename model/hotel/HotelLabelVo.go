package hotel

// HotelLabelVo 结构体
type HotelLabelVo struct {
	// 酒店标签文案
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// 酒店标签颜色
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// icon
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// refFieldName
	RefFieldName string `json:"ref_field_name,omitempty" xml:"ref_field_name,omitempty"`
	// reqFieldName
	ReqFieldName string `json:"req_field_name,omitempty" xml:"req_field_name,omitempty"`
	// 酒店标签id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// index
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// pos
	Pos int64 `json:"pos,omitempty" xml:"pos,omitempty"`
	// value
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// show
	Show bool `json:"show,omitempty" xml:"show,omitempty"`
}
