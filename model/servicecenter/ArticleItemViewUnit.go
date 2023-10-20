package servicecenter

import (
	"sync"
)

// ArticleItemViewUnit 结构体
type ArticleItemViewUnit struct {
	// 需要支付的价格，单位：元
	ActualPrice string `json:"actual_price,omitempty" xml:"actual_price,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误文案
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 收费项目code
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 收费项目名称
	ItemName string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 原价，单位：元
	OriginPrice string `json:"origin_price,omitempty" xml:"origin_price,omitempty"`
	// 优惠，单位：元
	PromPrice string `json:"prom_price,omitempty" xml:"prom_price,omitempty"`
	// 周期数，如1，3，6，12。对于周期型和周期计量型返回。
	CycNum int64 `json:"cyc_num,omitempty" xml:"cyc_num,omitempty"`
	// 1-年，2-月，3-日。对于周期型和周期计量型返回。
	CycUnit int64 `json:"cyc_unit,omitempty" xml:"cyc_unit,omitempty"`
	// 数量。对于周期计量型返回计量数。
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 用户是否可以购买
	CanSub bool `json:"can_sub,omitempty" xml:"can_sub,omitempty"`
}

var poolArticleItemViewUnit = sync.Pool{
	New: func() any {
		return new(ArticleItemViewUnit)
	},
}

// GetArticleItemViewUnit() 从对象池中获取ArticleItemViewUnit
func GetArticleItemViewUnit() *ArticleItemViewUnit {
	return poolArticleItemViewUnit.Get().(*ArticleItemViewUnit)
}

// ReleaseArticleItemViewUnit 释放ArticleItemViewUnit
func ReleaseArticleItemViewUnit(v *ArticleItemViewUnit) {
	v.ActualPrice = ""
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.ItemCode = ""
	v.ItemName = ""
	v.OriginPrice = ""
	v.PromPrice = ""
	v.CycNum = 0
	v.CycUnit = 0
	v.Quantity = 0
	v.CanSub = false
	poolArticleItemViewUnit.Put(v)
}
