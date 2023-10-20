package nrt

import (
	"sync"
)

// CouponTemplateDto 结构体
type CouponTemplateDto struct {
	// 圈选商品列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 圈选卖场列表
	MallIds []int64 `json:"mall_ids,omitempty" xml:"mall_ids>int64,omitempty"`
	// 圈选门店列表
	StoreIds []int64 `json:"store_ids,omitempty" xml:"store_ids>int64,omitempty"`
	// 渠道
	Channel string `json:"channel,omitempty" xml:"channel,omitempty"`
	// 券名称
	CouponName string `json:"coupon_name,omitempty" xml:"coupon_name,omitempty"`
	// 创建钉钉ID
	Creator string `json:"creator,omitempty" xml:"creator,omitempty"`
	// 修改钉钉ID
	ModifiedBy string `json:"modified_by,omitempty" xml:"modified_by,omitempty"`
	// 发放结束时间yyyy-MM-dd HH:mm:ss
	SendEndTime string `json:"send_end_time,omitempty" xml:"send_end_time,omitempty"`
	// 发放开始时间yyyy-MM-dd HH:mm:ss
	SendStartTime string `json:"send_start_time,omitempty" xml:"send_start_time,omitempty"`
	// 使用结束时间:yyyy-MM-dd HH:mm:ss
	UseEndTime string `json:"use_end_time,omitempty" xml:"use_end_time,omitempty"`
	// 使用开始时间:yyyy-MM-dd HH:mm:ss
	UseStartTime string `json:"use_start_time,omitempty" xml:"use_start_time,omitempty"`
	// 发券code
	Uuid string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// 券type
	CouponType int64 `json:"coupon_type,omitempty" xml:"coupon_type,omitempty"`
	// 面额
	Discount int64 `json:"discount,omitempty" xml:"discount,omitempty"`
	// 券模板ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 门槛
	StartFee int64 `json:"start_fee,omitempty" xml:"start_fee,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 领取后可以使用的相对时间，秒或天为单位
	UseTime int64 `json:"use_time,omitempty" xml:"use_time,omitempty"`
	// 使用时间类型
	UseTimeType int64 `json:"use_time_type,omitempty" xml:"use_time_type,omitempty"`
	// 版本号
	Version int64 `json:"version,omitempty" xml:"version,omitempty"`
	// 卖家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 券模板ID
	CouponTemplateId int64 `json:"coupon_template_id,omitempty" xml:"coupon_template_id,omitempty"`
}

var poolCouponTemplateDto = sync.Pool{
	New: func() any {
		return new(CouponTemplateDto)
	},
}

// GetCouponTemplateDto() 从对象池中获取CouponTemplateDto
func GetCouponTemplateDto() *CouponTemplateDto {
	return poolCouponTemplateDto.Get().(*CouponTemplateDto)
}

// ReleaseCouponTemplateDto 释放CouponTemplateDto
func ReleaseCouponTemplateDto(v *CouponTemplateDto) {
	v.ItemIds = v.ItemIds[:0]
	v.MallIds = v.MallIds[:0]
	v.StoreIds = v.StoreIds[:0]
	v.Channel = ""
	v.CouponName = ""
	v.Creator = ""
	v.ModifiedBy = ""
	v.SendEndTime = ""
	v.SendStartTime = ""
	v.UseEndTime = ""
	v.UseStartTime = ""
	v.Uuid = ""
	v.CouponType = 0
	v.Discount = 0
	v.Id = 0
	v.StartFee = 0
	v.Status = 0
	v.UseTime = 0
	v.UseTimeType = 0
	v.Version = 0
	v.SellerId = 0
	v.CouponTemplateId = 0
	poolCouponTemplateDto.Put(v)
}
