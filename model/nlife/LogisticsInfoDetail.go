package nlife

import (
	"sync"
)

// LogisticsInfoDetail 结构体
type LogisticsInfoDetail struct {
	// 商品列表：[“货码:数量”]，码可以是条形码(sku级别)也可以是零售加唯一码(货级别)、零售+ itemId+&#34;_&#34;+skuId，唯一码数量一定是1
	GoodsIds []string `json:"goods_ids,omitempty" xml:"goods_ids>string,omitempty"`
	// 物流单号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 物流公司名
	LogisticsCompany string `json:"logistics_company,omitempty" xml:"logistics_company,omitempty"`
	// 发货时间
	DeliverTime string `json:"deliver_time,omitempty" xml:"deliver_time,omitempty"`
}

var poolLogisticsInfoDetail = sync.Pool{
	New: func() any {
		return new(LogisticsInfoDetail)
	},
}

// GetLogisticsInfoDetail() 从对象池中获取LogisticsInfoDetail
func GetLogisticsInfoDetail() *LogisticsInfoDetail {
	return poolLogisticsInfoDetail.Get().(*LogisticsInfoDetail)
}

// ReleaseLogisticsInfoDetail 释放LogisticsInfoDetail
func ReleaseLogisticsInfoDetail(v *LogisticsInfoDetail) {
	v.GoodsIds = v.GoodsIds[:0]
	v.LogisticsNo = ""
	v.LogisticsCompany = ""
	v.DeliverTime = ""
	poolLogisticsInfoDetail.Put(v)
}
