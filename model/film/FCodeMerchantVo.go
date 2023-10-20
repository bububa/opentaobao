package film

import (
	"sync"
)

// FCodeMerchantVo 结构体
type FCodeMerchantVo struct {
	// 码过期时间
	GmtExpire string `json:"gmt_expire,omitempty" xml:"gmt_expire,omitempty"`
	// code
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 码生成任务ID
	GenTaskId int64 `json:"gen_task_id,omitempty" xml:"gen_task_id,omitempty"`
	// 码可抵用金额
	CostPrice int64 `json:"cost_price,omitempty" xml:"cost_price,omitempty"`
}

var poolFCodeMerchantVo = sync.Pool{
	New: func() any {
		return new(FCodeMerchantVo)
	},
}

// GetFCodeMerchantVo() 从对象池中获取FCodeMerchantVo
func GetFCodeMerchantVo() *FCodeMerchantVo {
	return poolFCodeMerchantVo.Get().(*FCodeMerchantVo)
}

// ReleaseFCodeMerchantVo 释放FCodeMerchantVo
func ReleaseFCodeMerchantVo(v *FCodeMerchantVo) {
	v.GmtExpire = ""
	v.Code = ""
	v.GenTaskId = 0
	v.CostPrice = 0
	poolFCodeMerchantVo.Put(v)
}
