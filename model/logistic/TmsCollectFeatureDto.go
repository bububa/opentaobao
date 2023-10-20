package logistic

import (
	"sync"
)

// TmsCollectFeatureDto 结构体
type TmsCollectFeatureDto struct {
	// 子运单号
	SubMailNo string `json:"sub_mail_no,omitempty" xml:"sub_mail_no,omitempty"`
}

var poolTmsCollectFeatureDto = sync.Pool{
	New: func() any {
		return new(TmsCollectFeatureDto)
	},
}

// GetTmsCollectFeatureDto() 从对象池中获取TmsCollectFeatureDto
func GetTmsCollectFeatureDto() *TmsCollectFeatureDto {
	return poolTmsCollectFeatureDto.Get().(*TmsCollectFeatureDto)
}

// ReleaseTmsCollectFeatureDto 释放TmsCollectFeatureDto
func ReleaseTmsCollectFeatureDto(v *TmsCollectFeatureDto) {
	v.SubMailNo = ""
	poolTmsCollectFeatureDto.Put(v)
}
