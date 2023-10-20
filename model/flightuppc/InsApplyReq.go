package flightuppc

import (
	"sync"
)

// InsApplyReq 结构体
type InsApplyReq struct {
	// 投保参数列表，通过险种聚类
	InsProductParams []InsProductBaseParam `json:"ins_product_params,omitempty" xml:"ins_product_params>ins_product_base_param,omitempty"`
}

var poolInsApplyReq = sync.Pool{
	New: func() any {
		return new(InsApplyReq)
	},
}

// GetInsApplyReq() 从对象池中获取InsApplyReq
func GetInsApplyReq() *InsApplyReq {
	return poolInsApplyReq.Get().(*InsApplyReq)
}

// ReleaseInsApplyReq 释放InsApplyReq
func ReleaseInsApplyReq(v *InsApplyReq) {
	v.InsProductParams = v.InsProductParams[:0]
	poolInsApplyReq.Put(v)
}
