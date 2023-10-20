package eleenterprisecartnew

import (
	"sync"
)

// EnterpriseData 结构体
type EnterpriseData struct {
	// 额外信息
	Extras []Extra `json:"extras,omitempty" xml:"extras>extra,omitempty"`
	// 购物车篮子
	Groups []string `json:"groups,omitempty" xml:"groups>string,omitempty"`
	// 购物车原始金额
	OriginalTotal string `json:"original_total,omitempty" xml:"original_total,omitempty"`
	// 购物车金额
	Total string `json:"total,omitempty" xml:"total,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty" xml:"phone,omitempty"`
	// 费用说明
	ServiceFeeExplanation string `json:"service_fee_explanation,omitempty" xml:"service_fee_explanation,omitempty"`
	// 起送价
	DeliverAmount string `json:"deliver_amount,omitempty" xml:"deliver_amount,omitempty"`
	// 购物车id
	Id string `json:"id,omitempty" xml:"id,omitempty"`
	// 创建购物车时间戳
	CreateTime int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
}

var poolEnterpriseData = sync.Pool{
	New: func() any {
		return new(EnterpriseData)
	},
}

// GetEnterpriseData() 从对象池中获取EnterpriseData
func GetEnterpriseData() *EnterpriseData {
	return poolEnterpriseData.Get().(*EnterpriseData)
}

// ReleaseEnterpriseData 释放EnterpriseData
func ReleaseEnterpriseData(v *EnterpriseData) {
	v.Extras = v.Extras[:0]
	v.Groups = v.Groups[:0]
	v.OriginalTotal = ""
	v.Total = ""
	v.Phone = ""
	v.ServiceFeeExplanation = ""
	v.DeliverAmount = ""
	v.Id = ""
	v.CreateTime = 0
	poolEnterpriseData.Put(v)
}
