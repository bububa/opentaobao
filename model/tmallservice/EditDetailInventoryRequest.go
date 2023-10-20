package tmallservice

import (
	"sync"
)

// EditDetailInventoryRequest 结构体
type EditDetailInventoryRequest struct {
	// 目标库存key。如果容量的时间间隔为天时，则必须为yyyy-MM-dd格式；如果容量的时间间隔为小时，则必须为yyyy-MM-dd HH:00-HH:00格式
	TargetInventoryKey string `json:"target_inventory_key,omitempty" xml:"target_inventory_key,omitempty"`
	// 不为0的整数
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
}

var poolEditDetailInventoryRequest = sync.Pool{
	New: func() any {
		return new(EditDetailInventoryRequest)
	},
}

// GetEditDetailInventoryRequest() 从对象池中获取EditDetailInventoryRequest
func GetEditDetailInventoryRequest() *EditDetailInventoryRequest {
	return poolEditDetailInventoryRequest.Get().(*EditDetailInventoryRequest)
}

// ReleaseEditDetailInventoryRequest 释放EditDetailInventoryRequest
func ReleaseEditDetailInventoryRequest(v *EditDetailInventoryRequest) {
	v.TargetInventoryKey = ""
	v.Quantity = 0
	poolEditDetailInventoryRequest.Put(v)
}
