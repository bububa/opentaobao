package tbk

import (
	"sync"
)

// BybtInfoDto 结构体
type BybtInfoDto struct {
	// 百亿补贴商品特征标签，eg.今日发货、晚发补偿、限购一件等
	BybtItemTags []string `json:"bybt_item_tags,omitempty" xml:"bybt_item_tags>string,omitempty"`
	// 百亿补贴品牌logo
	BybtBrandLogo string `json:"bybt_brand_logo,omitempty" xml:"bybt_brand_logo,omitempty"`
	// 百亿补贴白底图
	BybtPicUrl string `json:"bybt_pic_url,omitempty" xml:"bybt_pic_url,omitempty"`
	// 百亿补贴专属券面额，仅限百亿补贴场景透出
	BybtCouponAmount string `json:"bybt_coupon_amount,omitempty" xml:"bybt_coupon_amount,omitempty"`
	// 百亿补贴页面实时价
	BybtShowPrice string `json:"bybt_show_price,omitempty" xml:"bybt_show_price,omitempty"`
	// 全网对比参考价格
	BybtLowestPrice string `json:"bybt_lowest_price,omitempty" xml:"bybt_lowest_price,omitempty"`
	// 商品的百亿补贴开始时间
	BybtEndTime string `json:"bybt_end_time,omitempty" xml:"bybt_end_time,omitempty"`
	// 商品的百亿补贴结束时间
	BybtStartTime string `json:"bybt_start_time,omitempty" xml:"bybt_start_time,omitempty"`
}

var poolBybtInfoDto = sync.Pool{
	New: func() any {
		return new(BybtInfoDto)
	},
}

// GetBybtInfoDto() 从对象池中获取BybtInfoDto
func GetBybtInfoDto() *BybtInfoDto {
	return poolBybtInfoDto.Get().(*BybtInfoDto)
}

// ReleaseBybtInfoDto 释放BybtInfoDto
func ReleaseBybtInfoDto(v *BybtInfoDto) {
	v.BybtItemTags = v.BybtItemTags[:0]
	v.BybtBrandLogo = ""
	v.BybtPicUrl = ""
	v.BybtCouponAmount = ""
	v.BybtShowPrice = ""
	v.BybtLowestPrice = ""
	v.BybtEndTime = ""
	v.BybtStartTime = ""
	poolBybtInfoDto.Put(v)
}
