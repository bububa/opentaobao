package tmallservice

import (
	"sync"
)

// Spb2bOderQuery 结构体
type Spb2bOderQuery struct {
	// 开始日期
	StartDate string `json:"start_date,omitempty" xml:"start_date,omitempty"`
	// 结束日期
	EndDate string `json:"end_date,omitempty" xml:"end_date,omitempty"`
	// 售卖商昵称
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
	// 服务code
	ServiceCode string `json:"service_code,omitempty" xml:"service_code,omitempty"`
	// 服务状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 每页个数
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 最后条数
	EndRow int64 `json:"end_row,omitempty" xml:"end_row,omitempty"`
	// 开始条数
	StartRow int64 `json:"start_row,omitempty" xml:"start_row,omitempty"`
	// 售卖商id
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 是否分页
	NeedByPage bool `json:"need_by_page,omitempty" xml:"need_by_page,omitempty"`
	// 是否查询新供给ssc订购数据
	NewSupplySubscriberData bool `json:"new_supply_subscriber_data,omitempty" xml:"new_supply_subscriber_data,omitempty"`
}

var poolSpb2bOderQuery = sync.Pool{
	New: func() any {
		return new(Spb2bOderQuery)
	},
}

// GetSpb2bOderQuery() 从对象池中获取Spb2bOderQuery
func GetSpb2bOderQuery() *Spb2bOderQuery {
	return poolSpb2bOderQuery.Get().(*Spb2bOderQuery)
}

// ReleaseSpb2bOderQuery 释放Spb2bOderQuery
func ReleaseSpb2bOderQuery(v *Spb2bOderQuery) {
	v.StartDate = ""
	v.EndDate = ""
	v.SellerNick = ""
	v.ServiceCode = ""
	v.Status = 0
	v.PageSize = 0
	v.EndRow = 0
	v.StartRow = 0
	v.SellerId = 0
	v.CurrentPage = 0
	v.NeedByPage = false
	v.NewSupplySubscriberData = false
	poolSpb2bOderQuery.Put(v)
}
