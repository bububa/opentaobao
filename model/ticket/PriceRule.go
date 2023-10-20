package ticket

import (
	"sync"
)

// PriceRule 结构体
type PriceRule struct {
	// 每日价格库存
	DateInventorys []DateInventory `json:"date_inventorys,omitempty" xml:"date_inventorys>date_inventory,omitempty"`
	// 商户票种规则id
	OutRuleId string `json:"out_rule_id,omitempty" xml:"out_rule_id,omitempty"`
	// sku的商家编码（用于区分在 不同票种下使用同一outRuleId的情况）
	OutSkuId string `json:"out_sku_id,omitempty" xml:"out_sku_id,omitempty"`
	// 可选，1-全量更新，2-增量更新（增加或覆盖某一天的价格库存），不传默认为1。
	UploadType int64 `json:"upload_type,omitempty" xml:"upload_type,omitempty"`
}

var poolPriceRule = sync.Pool{
	New: func() any {
		return new(PriceRule)
	},
}

// GetPriceRule() 从对象池中获取PriceRule
func GetPriceRule() *PriceRule {
	return poolPriceRule.Get().(*PriceRule)
}

// ReleasePriceRule 释放PriceRule
func ReleasePriceRule(v *PriceRule) {
	v.DateInventorys = v.DateInventorys[:0]
	v.OutRuleId = ""
	v.OutSkuId = ""
	v.UploadType = 0
	poolPriceRule.Put(v)
}
