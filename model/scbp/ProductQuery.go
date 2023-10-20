package scbp

import (
	"sync"
)

// ProductQuery 结构体
type ProductQuery struct {
	// 结束时间 当inteval=7或30的时候 不需要填写
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 排序字段，目前支持传入下面5个值中的其中一个，不传的话默认使用impression_cnt，括号中为字段含义解释 impression_cnt (曝光量)，click_cost_avg (平均点击花费)，click_cnt(点击量)cost(总花费)ctr(点击率)
	OrderStr string `json:"order_str,omitempty" xml:"order_str,omitempty"`
	// 开始时间 当inteval=7或30的时候 不需要填写
	BeginDate string `json:"begin_date,omitempty" xml:"begin_date,omitempty"`
	// 区间 只能为1 7 30
	Inteval int64 `json:"inteval,omitempty" xml:"inteval,omitempty"`
	// 每页行数
	PerPageSize int64 `json:"per_page_size,omitempty" xml:"per_page_size,omitempty"`
	// 第几页
	ToPage int64 `json:"to_page,omitempty" xml:"to_page,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

var poolProductQuery = sync.Pool{
	New: func() any {
		return new(ProductQuery)
	},
}

// GetProductQuery() 从对象池中获取ProductQuery
func GetProductQuery() *ProductQuery {
	return poolProductQuery.Get().(*ProductQuery)
}

// ReleaseProductQuery 释放ProductQuery
func ReleaseProductQuery(v *ProductQuery) {
	v.EndDate = ""
	v.OrderStr = ""
	v.BeginDate = ""
	v.Inteval = 0
	v.PerPageSize = 0
	v.ToPage = 0
	v.ProductId = 0
	poolProductQuery.Put(v)
}
