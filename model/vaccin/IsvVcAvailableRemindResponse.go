package vaccin

import (
	"sync"
)

// IsvVcAvailableRemindResponse 结构体
type IsvVcAvailableRemindResponse struct {
	// 登记单主键id。total如果大于200，则只显示200个id
	OuterRegisterIdList []string `json:"outer_register_id_list,omitempty" xml:"outer_register_id_list>string,omitempty"`
	// 登记单总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolIsvVcAvailableRemindResponse = sync.Pool{
	New: func() any {
		return new(IsvVcAvailableRemindResponse)
	},
}

// GetIsvVcAvailableRemindResponse() 从对象池中获取IsvVcAvailableRemindResponse
func GetIsvVcAvailableRemindResponse() *IsvVcAvailableRemindResponse {
	return poolIsvVcAvailableRemindResponse.Get().(*IsvVcAvailableRemindResponse)
}

// ReleaseIsvVcAvailableRemindResponse 释放IsvVcAvailableRemindResponse
func ReleaseIsvVcAvailableRemindResponse(v *IsvVcAvailableRemindResponse) {
	v.OuterRegisterIdList = v.OuterRegisterIdList[:0]
	v.Total = 0
	poolIsvVcAvailableRemindResponse.Put(v)
}
