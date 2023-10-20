package xhotelitem

import (
	"sync"
)

// SellerStatExposureElementList 结构体
type SellerStatExposureElementList struct {
	// 日期
	StatDate string `json:"stat_date,omitempty" xml:"stat_date,omitempty"`
	// 曝光率
	ExposedPercent string `json:"exposed_percent,omitempty" xml:"exposed_percent,omitempty"`
	// shid维度访问量
	ShidTotalAmount string `json:"shid_total_amount,omitempty" xml:"shid_total_amount,omitempty"`
	// 对应商品曝光数量
	ExposedAmount string `json:"exposed_amount,omitempty" xml:"exposed_amount,omitempty"`
}

var poolSellerStatExposureElementList = sync.Pool{
	New: func() any {
		return new(SellerStatExposureElementList)
	},
}

// GetSellerStatExposureElementList() 从对象池中获取SellerStatExposureElementList
func GetSellerStatExposureElementList() *SellerStatExposureElementList {
	return poolSellerStatExposureElementList.Get().(*SellerStatExposureElementList)
}

// ReleaseSellerStatExposureElementList 释放SellerStatExposureElementList
func ReleaseSellerStatExposureElementList(v *SellerStatExposureElementList) {
	v.StatDate = ""
	v.ExposedPercent = ""
	v.ShidTotalAmount = ""
	v.ExposedAmount = ""
	poolSellerStatExposureElementList.Put(v)
}
