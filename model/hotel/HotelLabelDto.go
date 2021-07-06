package hotel

// HotelLabelDto 结构体
type HotelLabelDto struct {
	// 颜色
	Color string `json:"color,omitempty" xml:"color,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 图片
	Icon string `json:"icon,omitempty" xml:"icon,omitempty"`
	// 需要映射到PC的参数（拼装生成pc请求值）
	RefFieldName string `json:"ref_field_name,omitempty" xml:"ref_field_name,omitempty"`
	// 需要映射到PC的参数(拼装生成pc返回值)
	ReqFieldName string `json:"req_field_name,omitempty" xml:"req_field_name,omitempty"`
	// 内容
	Text string `json:"text,omitempty" xml:"text,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 索引
	Index int64 `json:"index,omitempty" xml:"index,omitempty"`
	// 位置
	Pos int64 `json:"pos,omitempty" xml:"pos,omitempty"`
	// 值
	Value int64 `json:"value,omitempty" xml:"value,omitempty"`
	// is show label
	Show bool `json:"show,omitempty" xml:"show,omitempty"`
}
