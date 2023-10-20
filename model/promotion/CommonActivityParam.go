package promotion

import (
	"sync"
)

// CommonActivityParam 结构体
type CommonActivityParam struct {
	// 商家优惠券活动id
	OutActId string `json:"out_act_id,omitempty" xml:"out_act_id,omitempty"`
	// 五道口优惠券活动id
	ActivityId int64 `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
}

var poolCommonActivityParam = sync.Pool{
	New: func() any {
		return new(CommonActivityParam)
	},
}

// GetCommonActivityParam() 从对象池中获取CommonActivityParam
func GetCommonActivityParam() *CommonActivityParam {
	return poolCommonActivityParam.Get().(*CommonActivityParam)
}

// ReleaseCommonActivityParam 释放CommonActivityParam
func ReleaseCommonActivityParam(v *CommonActivityParam) {
	v.OutActId = ""
	v.ActivityId = 0
	poolCommonActivityParam.Put(v)
}
