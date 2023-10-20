package wdk

import (
	"sync"
)

// SkuActivityElementDto 结构体
type SkuActivityElementDto struct {
	// 商品渠道配置信息
	SkuChannelConfigs []SkuChannelConfigDto `json:"sku_channel_configs,omitempty" xml:"sku_channel_configs>sku_channel_config_dto,omitempty"`
	// 商品编码，与商品条码必选其一，或者同时传入以商品条码为准
	SkuCode string `json:"sku_code,omitempty" xml:"sku_code,omitempty"`
	// 赠品编码，不填默认赠品为商品本身，赠品条码，不填默认赠品为商品本身，若主品传入的是编码，则赠品需传入编码或者不传
	GiftSkuCode string `json:"gift_sku_code,omitempty" xml:"gift_sku_code,omitempty"`
	// 商品条码，与商品编码必选其一，或者同时传入以商品条码为准
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 赠品条码，不填默认赠品为商品本身，若主品传入的是条码，则赠品需传入条码或者不传
	GiftBarcode string `json:"gift_barcode,omitempty" xml:"gift_barcode,omitempty"`
	// 操作人ID
	CreatorId string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// 操作人姓名
	CreatorName string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// 商品条码
	BarCode string `json:"bar_code,omitempty" xml:"bar_code,omitempty"`
	// 限购配置信息
	Limit *LimitDto `json:"limit,omitempty" xml:"limit,omitempty"`
	// 买N赠M的M参数，赠多少件赠品（目前仅支持买N赠1，giftNum默认为1）
	GiftNum int64 `json:"gift_num,omitempty" xml:"gift_num,omitempty"`
	// 买N赠M的N参数，买多少件可赠
	BuyNum int64 `json:"buy_num,omitempty" xml:"buy_num,omitempty"`
	// 一口价金额【分】
	FixPriceMoney int64 `json:"fix_price_money,omitempty" xml:"fix_price_money,omitempty"`
	// 打折金额【1000位底数】，900代表9折
	DiscountRate int64 `json:"discount_rate,omitempty" xml:"discount_rate,omitempty"`
	// 减钱金额【分】
	DecreaseMoney int64 `json:"decrease_money,omitempty" xml:"decrease_money,omitempty"`
	// 营销活动ID
	ActId int64 `json:"act_id,omitempty" xml:"act_id,omitempty"`
}

var poolSkuActivityElementDto = sync.Pool{
	New: func() any {
		return new(SkuActivityElementDto)
	},
}

// GetSkuActivityElementDto() 从对象池中获取SkuActivityElementDto
func GetSkuActivityElementDto() *SkuActivityElementDto {
	return poolSkuActivityElementDto.Get().(*SkuActivityElementDto)
}

// ReleaseSkuActivityElementDto 释放SkuActivityElementDto
func ReleaseSkuActivityElementDto(v *SkuActivityElementDto) {
	v.SkuChannelConfigs = v.SkuChannelConfigs[:0]
	v.SkuCode = ""
	v.GiftSkuCode = ""
	v.Barcode = ""
	v.GiftBarcode = ""
	v.CreatorId = ""
	v.CreatorName = ""
	v.BarCode = ""
	v.Limit = nil
	v.GiftNum = 0
	v.BuyNum = 0
	v.FixPriceMoney = 0
	v.DiscountRate = 0
	v.DecreaseMoney = 0
	v.ActId = 0
	poolSkuActivityElementDto.Put(v)
}
