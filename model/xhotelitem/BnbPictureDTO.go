package xhotelitem

// BnbPictureDto 结构体
type BnbPictureDto struct {
	// 图片地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// type表示图片类型，取值范围只能是：[周边, 外观, 商务中心, 健身房, 其他, 会议室, 餐厅, 浴室, 客房, 公共区域, 娱乐设施, 大堂, 泳池]
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// 图片属性 取值范围只能是：[普通图, 平面图, 全景图]
	Attribute string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// 图片描述
	Des string `json:"des,omitempty" xml:"des,omitempty"`
	// 是否主图  主图只能有一个，如果有多个或者没有，则会报错
	Ismain bool `json:"ismain,omitempty" xml:"ismain,omitempty"`
}
