package alihealth2

import (
	"sync"
)

// AlibabaAlihealthDentalItemListMtopResult 结构体
type AlibabaAlihealthDentalItemListMtopResult struct {
	// model
	Goods []NormalGoodsVo `json:"goods,omitempty" xml:"goods>normal_goods_vo,omitempty"`
	// 200
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 成功
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// true
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDentalItemListMtopResult = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDentalItemListMtopResult)
	},
}

// GetAlibabaAlihealthDentalItemListMtopResult() 从对象池中获取AlibabaAlihealthDentalItemListMtopResult
func GetAlibabaAlihealthDentalItemListMtopResult() *AlibabaAlihealthDentalItemListMtopResult {
	return poolAlibabaAlihealthDentalItemListMtopResult.Get().(*AlibabaAlihealthDentalItemListMtopResult)
}

// ReleaseAlibabaAlihealthDentalItemListMtopResult 释放AlibabaAlihealthDentalItemListMtopResult
func ReleaseAlibabaAlihealthDentalItemListMtopResult(v *AlibabaAlihealthDentalItemListMtopResult) {
	v.Goods = v.Goods[:0]
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Success = false
	poolAlibabaAlihealthDentalItemListMtopResult.Put(v)
}
