package alihouse

import (
	"sync"
)

// BusinessActivitySignatureDto 结构体
type BusinessActivitySignatureDto struct {
	// 下旋文
	BackspinText string `json:"backspin_text,omitempty" xml:"backspin_text,omitempty"`
	// 横向文
	HorizontalText string `json:"horizontal_text,omitempty" xml:"horizontal_text,omitempty"`
	// 签章名称
	SignatureName string `json:"signature_name,omitempty" xml:"signature_name,omitempty"`
	// 签章图片地址
	SealUrl string `json:"seal_url,omitempty" xml:"seal_url,omitempty"`
	// 中心图案类型
	CenterPattern int64 `json:"center_pattern,omitempty" xml:"center_pattern,omitempty"`
	// 签章颜色
	SignatureColor int64 `json:"signature_color,omitempty" xml:"signature_color,omitempty"`
	// 签章类型
	SignatureType int64 `json:"signature_type,omitempty" xml:"signature_type,omitempty"`
	// ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolBusinessActivitySignatureDto = sync.Pool{
	New: func() any {
		return new(BusinessActivitySignatureDto)
	},
}

// GetBusinessActivitySignatureDto() 从对象池中获取BusinessActivitySignatureDto
func GetBusinessActivitySignatureDto() *BusinessActivitySignatureDto {
	return poolBusinessActivitySignatureDto.Get().(*BusinessActivitySignatureDto)
}

// ReleaseBusinessActivitySignatureDto 释放BusinessActivitySignatureDto
func ReleaseBusinessActivitySignatureDto(v *BusinessActivitySignatureDto) {
	v.BackspinText = ""
	v.HorizontalText = ""
	v.SignatureName = ""
	v.SealUrl = ""
	v.CenterPattern = 0
	v.SignatureColor = 0
	v.SignatureType = 0
	v.Id = 0
	poolBusinessActivitySignatureDto.Put(v)
}
