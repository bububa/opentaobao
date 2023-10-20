package xhotelonlineorder

import (
	"sync"
)

// UpdateRateParam 结构体
type UpdateRateParam struct {
	// rate更新列表
	UpdateRateDOList []UpdateRateDo `json:"update_rate_d_o_list,omitempty" xml:"update_rate_d_o_list>update_rate_do,omitempty"`
	// 过期时间
	ExpireTime string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// 供应商
	Supplier string `json:"supplier,omitempty" xml:"supplier,omitempty"`
}

var poolUpdateRateParam = sync.Pool{
	New: func() any {
		return new(UpdateRateParam)
	},
}

// GetUpdateRateParam() 从对象池中获取UpdateRateParam
func GetUpdateRateParam() *UpdateRateParam {
	return poolUpdateRateParam.Get().(*UpdateRateParam)
}

// ReleaseUpdateRateParam 释放UpdateRateParam
func ReleaseUpdateRateParam(v *UpdateRateParam) {
	v.UpdateRateDOList = v.UpdateRateDOList[:0]
	v.ExpireTime = ""
	v.Supplier = ""
	poolUpdateRateParam.Put(v)
}
