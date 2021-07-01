package damai

// ThirdFaceElementPushOpenParam 结构体
type ThirdFaceElementPushOpenParam struct {
	// 变量类型
	DataType int64 `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 日期格式
	DateFormat string `json:"date_format,omitempty" xml:"date_format,omitempty"`
	// 元素类型
	ElementType int64 `json:"element_type,omitempty" xml:"element_type,omitempty"`
	// 扩展字段名称
	ExtName string `json:"ext_name,omitempty" xml:"ext_name,omitempty"`
	// 扩展字段类型
	ExtType string `json:"ext_type,omitempty" xml:"ext_type,omitempty"`
	// 票面id
	FaceId int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
	// 字体
	Font string `json:"font,omitempty" xml:"font,omitempty"`
	// 字体朝向
	FontOrientation int64 `json:"font_orientation,omitempty" xml:"font_orientation,omitempty"`
	// 字体形状
	FontShape int64 `json:"font_shape,omitempty" xml:"font_shape,omitempty"`
	// 字号
	FontSize int64 `json:"font_size,omitempty" xml:"font_size,omitempty"`
	// 元素高度
	Height int64 `json:"height,omitempty" xml:"height,omitempty"`
	// 横坐标
	HorizontalCoordinate int64 `json:"horizontal_coordinate,omitempty" xml:"horizontal_coordinate,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 图片url
	StaticPicUrl string `json:"static_pic_url,omitempty" xml:"static_pic_url,omitempty"`
	// 图片内容
	StaticTextContent string `json:"static_text_content,omitempty" xml:"static_text_content,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 纵坐标
	VerticalCoordinate int64 `json:"vertical_coordinate,omitempty" xml:"vertical_coordinate,omitempty"`
	// 元素宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}
