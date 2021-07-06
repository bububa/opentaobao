package alitripmerchant

// HotelPictureDto 结构体
type HotelPictureDto struct {
	// 图片集合
	PictureAddress []string `json:"picture_address,omitempty" xml:"picture_address>string,omitempty"`
	// 类型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 类型编码
	TypeCode string `json:"type_code,omitempty" xml:"type_code,omitempty"`
}
