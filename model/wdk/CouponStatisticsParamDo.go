package wdk

import (
	"sync"
)

// CouponStatisticsParamDo 结构体
type CouponStatisticsParamDo struct {
	// 品牌名称数组
	BrandNames []string `json:"brand_names,omitempty" xml:"brand_names>string,omitempty"`
	// 日期，格式为yyyymmdd
	StatisticsDate string `json:"statistics_date,omitempty" xml:"statistics_date,omitempty"`
	// 页码，即当前第几页
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页记录数，不能超过200
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}

var poolCouponStatisticsParamDo = sync.Pool{
	New: func() any {
		return new(CouponStatisticsParamDo)
	},
}

// GetCouponStatisticsParamDo() 从对象池中获取CouponStatisticsParamDo
func GetCouponStatisticsParamDo() *CouponStatisticsParamDo {
	return poolCouponStatisticsParamDo.Get().(*CouponStatisticsParamDo)
}

// ReleaseCouponStatisticsParamDo 释放CouponStatisticsParamDo
func ReleaseCouponStatisticsParamDo(v *CouponStatisticsParamDo) {
	v.BrandNames = v.BrandNames[:0]
	v.StatisticsDate = ""
	v.PageIndex = 0
	v.PageSize = 0
	poolCouponStatisticsParamDo.Put(v)
}
