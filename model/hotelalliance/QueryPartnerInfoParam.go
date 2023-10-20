package hotelalliance

import (
	"sync"
)

// QueryPartnerInfoParam 结构体
type QueryPartnerInfoParam struct {
	// 合作商Id
	HocPartnerId int64 `json:"hoc_partner_id,omitempty" xml:"hoc_partner_id,omitempty"`
}

var poolQueryPartnerInfoParam = sync.Pool{
	New: func() any {
		return new(QueryPartnerInfoParam)
	},
}

// GetQueryPartnerInfoParam() 从对象池中获取QueryPartnerInfoParam
func GetQueryPartnerInfoParam() *QueryPartnerInfoParam {
	return poolQueryPartnerInfoParam.Get().(*QueryPartnerInfoParam)
}

// ReleaseQueryPartnerInfoParam 释放QueryPartnerInfoParam
func ReleaseQueryPartnerInfoParam(v *QueryPartnerInfoParam) {
	v.HocPartnerId = 0
	poolQueryPartnerInfoParam.Put(v)
}
