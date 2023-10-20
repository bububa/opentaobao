package traveltrade

import (
	"sync"
)

// HotelTicketVoucherDto 结构体
type HotelTicketVoucherDto struct {
	// 凭证码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 证件号
	CertificateId string `json:"certificate_id,omitempty" xml:"certificate_id,omitempty"`
	// 二维码图片链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 每张票或券可使用次数（如针对一码多刷，往返索道3张票1个码，每张票可使用次数为2，则该码可以刷6次
	AvailableNums int64 `json:"available_nums,omitempty" xml:"available_nums,omitempty"`
	// 凭证类型 1：票码， 2：券码
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
	// 已使用次数
	UsageNums int64 `json:"usage_nums,omitempty" xml:"usage_nums,omitempty"`
	// 业务类型：1：门票， 2：酒店
	BizType int64 `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
	// 凭证 可用/不可用
	CanUse bool `json:"can_use,omitempty" xml:"can_use,omitempty"`
}

var poolHotelTicketVoucherDto = sync.Pool{
	New: func() any {
		return new(HotelTicketVoucherDto)
	},
}

// GetHotelTicketVoucherDto() 从对象池中获取HotelTicketVoucherDto
func GetHotelTicketVoucherDto() *HotelTicketVoucherDto {
	return poolHotelTicketVoucherDto.Get().(*HotelTicketVoucherDto)
}

// ReleaseHotelTicketVoucherDto 释放HotelTicketVoucherDto
func ReleaseHotelTicketVoucherDto(v *HotelTicketVoucherDto) {
	v.Code = ""
	v.CertificateId = ""
	v.Url = ""
	v.AvailableNums = 0
	v.Type = 0
	v.UsageNums = 0
	v.BizType = 0
	v.CanUse = false
	poolHotelTicketVoucherDto.Put(v)
}
