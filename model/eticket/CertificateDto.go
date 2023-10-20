package eticket

import (
	"sync"
)

// CertificateDto 结构体
type CertificateDto struct {
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// endTime
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// outerId
	OuterId string `json:"outer_id,omitempty" xml:"outer_id,omitempty"`
	// qrCodeUrl
	QrCodeUrl string `json:"qr_code_url,omitempty" xml:"qr_code_url,omitempty"`
	// startTime
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// attributes
	Attributes *Attributes `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// availableNum
	AvailableNum int64 `json:"available_num,omitempty" xml:"available_num,omitempty"`
	// bizType
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// codeStatus
	CodeStatus int64 `json:"code_status,omitempty" xml:"code_status,omitempty"`
	// initialNum
	InitialNum int64 `json:"initial_num,omitempty" xml:"initial_num,omitempty"`
	// lockedNum
	LockedNum int64 `json:"locked_num,omitempty" xml:"locked_num,omitempty"`
	// usedNum
	UsedNum int64 `json:"used_num,omitempty" xml:"used_num,omitempty"`
}

var poolCertificateDto = sync.Pool{
	New: func() any {
		return new(CertificateDto)
	},
}

// GetCertificateDto() 从对象池中获取CertificateDto
func GetCertificateDto() *CertificateDto {
	return poolCertificateDto.Get().(*CertificateDto)
}

// ReleaseCertificateDto 释放CertificateDto
func ReleaseCertificateDto(v *CertificateDto) {
	v.Code = ""
	v.EndTime = ""
	v.OuterId = ""
	v.QrCodeUrl = ""
	v.StartTime = ""
	v.Attributes = nil
	v.AvailableNum = 0
	v.BizType = 0
	v.CodeStatus = 0
	v.InitialNum = 0
	v.LockedNum = 0
	v.UsedNum = 0
	poolCertificateDto.Put(v)
}
