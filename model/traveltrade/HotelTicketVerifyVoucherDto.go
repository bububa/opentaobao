package traveltrade

import (
	"sync"
)

// HotelTicketVerifyVoucherDto 结构体
type HotelTicketVerifyVoucherDto struct {
	// 凭证码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 使用时间：yyyy-MM-dd HH:mm:ss
	UseDate string `json:"use_date,omitempty" xml:"use_date,omitempty"`
	// 证件号
	CertificateId string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	// 凭证类型 1：票码， 2：券码
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 票或券 核销使用数量
	UsageNums int64 `json:"usage_nums,omitempty" xml:"usage_nums,omitempty"`
	// 业务类型：1：门票， 2：酒店
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolHotelTicketVerifyVoucherDto = sync.Pool{
	New: func() any {
		return new(HotelTicketVerifyVoucherDto)
	},
}

// GetHotelTicketVerifyVoucherDto() 从对象池中获取HotelTicketVerifyVoucherDto
func GetHotelTicketVerifyVoucherDto() *HotelTicketVerifyVoucherDto {
	return poolHotelTicketVerifyVoucherDto.Get().(*HotelTicketVerifyVoucherDto)
}

// ReleaseHotelTicketVerifyVoucherDto 释放HotelTicketVerifyVoucherDto
func ReleaseHotelTicketVerifyVoucherDto(v *HotelTicketVerifyVoucherDto) {
	v.Code = ""
	v.UseDate = ""
	v.CertificateId = ""
	v.Type = 0
	v.UsageNums = 0
	v.BizType = 0
	poolHotelTicketVerifyVoucherDto.Put(v)
}
