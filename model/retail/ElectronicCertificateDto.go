package retail

import (
	"sync"
)

// ElectronicCertificateDto 结构体
type ElectronicCertificateDto struct {
	// itemId
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolElectronicCertificateDto = sync.Pool{
	New: func() any {
		return new(ElectronicCertificateDto)
	},
}

// GetElectronicCertificateDto() 从对象池中获取ElectronicCertificateDto
func GetElectronicCertificateDto() *ElectronicCertificateDto {
	return poolElectronicCertificateDto.Get().(*ElectronicCertificateDto)
}

// ReleaseElectronicCertificateDto 释放ElectronicCertificateDto
func ReleaseElectronicCertificateDto(v *ElectronicCertificateDto) {
	v.ItemId = 0
	poolElectronicCertificateDto.Put(v)
}
