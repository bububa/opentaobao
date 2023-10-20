package wdk

import (
	"sync"
)

// ReturnBoxContainerRequest 结构体
type ReturnBoxContainerRequest struct {
	// 周转箱列表
	BoxCodeList []string `json:"box_code_list,omitempty" xml:"box_code_list>string,omitempty"`
	// 收箱交接单号
	HandOverNo string `json:"hand_over_no,omitempty" xml:"hand_over_no,omitempty"`
	// 仓编号
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 操作时间
	OperateTime string `json:"operate_time,omitempty" xml:"operate_time,omitempty"`
}

var poolReturnBoxContainerRequest = sync.Pool{
	New: func() any {
		return new(ReturnBoxContainerRequest)
	},
}

// GetReturnBoxContainerRequest() 从对象池中获取ReturnBoxContainerRequest
func GetReturnBoxContainerRequest() *ReturnBoxContainerRequest {
	return poolReturnBoxContainerRequest.Get().(*ReturnBoxContainerRequest)
}

// ReleaseReturnBoxContainerRequest 释放ReturnBoxContainerRequest
func ReleaseReturnBoxContainerRequest(v *ReturnBoxContainerRequest) {
	v.BoxCodeList = v.BoxCodeList[:0]
	v.HandOverNo = ""
	v.WarehouseCode = ""
	v.OperateTime = ""
	poolReturnBoxContainerRequest.Put(v)
}
