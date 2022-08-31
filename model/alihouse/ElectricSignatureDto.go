package alihouse

// ElectricSignatureDto 结构体
type ElectricSignatureDto struct {
	// 外部电子印章ID
	OuterSignatureId string `json:"outer_signature_id,omitempty" xml:"outer_signature_id,omitempty"`
	// 电子印章名称
	SignatureName string `json:"signature_name,omitempty" xml:"signature_name,omitempty"`
	// 横向文
	HorizontalText string `json:"horizontal_text,omitempty" xml:"horizontal_text,omitempty"`
	// 下弦文
	BackspinText string `json:"backspin_text,omitempty" xml:"backspin_text,omitempty"`
	// 电子签章形状(1-圆章 2-椭圆章)
	SignatureType int64 `json:"signature_type,omitempty" xml:"signature_type,omitempty"`
	// 电子签章颜色(1-红色 2-蓝色 3-黑色)
	SignatureColor int64 `json:"signature_color,omitempty" xml:"signature_color,omitempty"`
	// 中心图案类型(1-圆形有五角星 2-圆形无五角星)
	CenterPattern int64 `json:"center_pattern,omitempty" xml:"center_pattern,omitempty"`
	// 是否生效
	IsValid int64 `json:"is_valid,omitempty" xml:"is_valid,omitempty"`
	// 经营主体ID
	MerchantOpenId int64 `json:"merchant_open_id,omitempty" xml:"merchant_open_id,omitempty"`
}
