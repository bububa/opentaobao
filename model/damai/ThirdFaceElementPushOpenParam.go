package damai

import (
	"sync"
)

// ThirdFaceElementPushOpenParam 结构体
type ThirdFaceElementPushOpenParam struct {
	// 日期格式
	DateFormat string `json:"date_format,omitempty" xml:"date_format,omitempty"`
	// 扩展字段名称
	ExtName string `json:"ext_name,omitempty" xml:"ext_name,omitempty"`
	// 扩展字段类型
	ExtType string `json:"ext_type,omitempty" xml:"ext_type,omitempty"`
	// 字体
	Font string `json:"font,omitempty" xml:"font,omitempty"`
	// 推送时间
	PushTime string `json:"push_time,omitempty" xml:"push_time,omitempty"`
	// 图片url
	StaticPicUrl string `json:"static_pic_url,omitempty" xml:"static_pic_url,omitempty"`
	// 图片内容
	StaticTextContent string `json:"static_text_content,omitempty" xml:"static_text_content,omitempty"`
	// 商户密钥
	SupplierSecret string `json:"supplier_secret,omitempty" xml:"supplier_secret,omitempty"`
	// 变量类型
	DataType int64 `json:"data_type,omitempty" xml:"data_type,omitempty"`
	// 元素类型
	ElementType int64 `json:"element_type,omitempty" xml:"element_type,omitempty"`
	// 票面id
	FaceId int64 `json:"face_id,omitempty" xml:"face_id,omitempty"`
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
	// 系统id
	SystemId int64 `json:"system_id,omitempty" xml:"system_id,omitempty"`
	// 纵坐标
	VerticalCoordinate int64 `json:"vertical_coordinate,omitempty" xml:"vertical_coordinate,omitempty"`
	// 元素宽度
	Width int64 `json:"width,omitempty" xml:"width,omitempty"`
}

var poolThirdFaceElementPushOpenParam = sync.Pool{
	New: func() any {
		return new(ThirdFaceElementPushOpenParam)
	},
}

// GetThirdFaceElementPushOpenParam() 从对象池中获取ThirdFaceElementPushOpenParam
func GetThirdFaceElementPushOpenParam() *ThirdFaceElementPushOpenParam {
	return poolThirdFaceElementPushOpenParam.Get().(*ThirdFaceElementPushOpenParam)
}

// ReleaseThirdFaceElementPushOpenParam 释放ThirdFaceElementPushOpenParam
func ReleaseThirdFaceElementPushOpenParam(v *ThirdFaceElementPushOpenParam) {
	v.DateFormat = ""
	v.ExtName = ""
	v.ExtType = ""
	v.Font = ""
	v.PushTime = ""
	v.StaticPicUrl = ""
	v.StaticTextContent = ""
	v.SupplierSecret = ""
	v.DataType = 0
	v.ElementType = 0
	v.FaceId = 0
	v.FontOrientation = 0
	v.FontShape = 0
	v.FontSize = 0
	v.Height = 0
	v.HorizontalCoordinate = 0
	v.SystemId = 0
	v.VerticalCoordinate = 0
	v.Width = 0
	poolThirdFaceElementPushOpenParam.Put(v)
}
