package alink

// OuterDataBean 结构体
type OuterDataBean struct {
	// 上报内容，可以是数字，或者json格式文本
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 产品型号，在入驻时生成的
	Model string `json:"model,omitempty" xml:"model,omitempty"`
	// 数据类型，detail：明细数据，stat：统计数据
	DataType string `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 产品型号key，和model对应，在注册产品型号时分配
	ModelKey string `json:"model_key,omitempty" xml:"model_key,omitempty"`
	// 数据产生的时间
	StatTime string `json:"stat_time,omitempty" xml:"stat_time,omitempty"`
	// 外部数据的唯一id，比如：设备mac、sn等表示设备唯一性的数据，或者表示该统计数据的自定义id
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
}
