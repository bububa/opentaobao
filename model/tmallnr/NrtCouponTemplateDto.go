package tmallnr

import (
	"sync"
)

// NrtCouponTemplateDto 结构体
type NrtCouponTemplateDto struct {
	// 券渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 创建人
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改人
	ModifiedBy string `json:"modified_by,omitempty" xml:"modified_by,omitempty"`
	// 使用开始时间:yyyy-MM-dd HH:mm:ss
	StartTime string `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 使用结束时间:yyyy-MM-dd HH:mm:ss
	EndTime string `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 发放开始时间yyyy-MM-dd HH:mm:ss
	SendStartTime string `json:"send_start_time,omitempty" xml:"send_start_time,omitempty"`
	// 发放结束时间yyyy-MM-dd HH:mm:ss
	SendEndTime string `json:"send_end_time,omitempty" xml:"send_end_time,omitempty"`
	// 发放code
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 扩展信息
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 版本
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 券总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 券剩余数量
	ReserveCount int64 `json:"reserve_count,omitempty" xml:"reserve_count,omitempty"`
	// 门槛
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 面额
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 使用时间类型use_time_type
	UseTimeType int64 `json:"use_time_type,omitempty" xml:"use_time_type,omitempty"`
	// 领取后可以使用的相对时间，秒或天为单位
	UseTime int64 `json:"use_time,omitempty" xml:"use_time,omitempty"`
	// 券状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 券类型
	CouponType int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 券模版id
	CouponTemplateId int64 `json:"coupon_template_id,omitempty" xml:"coupon_template_id,omitempty"`
}

var poolNrtCouponTemplateDto = sync.Pool{
	New: func() any {
		return new(NrtCouponTemplateDto)
	},
}

// GetNrtCouponTemplateDto() 从对象池中获取NrtCouponTemplateDto
func GetNrtCouponTemplateDto() *NrtCouponTemplateDto {
	return poolNrtCouponTemplateDto.Get().(*NrtCouponTemplateDto)
}

// ReleaseNrtCouponTemplateDto 释放NrtCouponTemplateDto
func ReleaseNrtCouponTemplateDto(v *NrtCouponTemplateDto) {
	v.Channel = ""
	v.Creator = ""
	v.ModifiedBy = ""
	v.StartTime = ""
	v.EndTime = ""
	v.CouponName = ""
	v.SendStartTime = ""
	v.SendEndTime = ""
	v.Uuid = ""
	v.Extra = ""
	v.Version = 0
	v.TotalCount = 0
	v.ReserveCount = 0
	v.StartFee = 0
	v.Discount = 0
	v.UseTimeType = 0
	v.UseTime = 0
	v.Status = 0
	v.CouponType = 0
	v.CouponTemplateId = 0
	poolNrtCouponTemplateDto.Put(v)
}
