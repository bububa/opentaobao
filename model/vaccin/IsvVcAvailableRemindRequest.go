package vaccin

import (
	"sync"
)

// IsvVcAvailableRemindRequest 结构体
type IsvVcAvailableRemindRequest struct {
	// 阿里健康的疫苗id
	VaccineId string `json:"vaccine_id,omitempty" xml:"vaccine_id,omitempty"`
	// 阿里健康的pov id
	PovId string `json:"pov_id,omitempty" xml:"pov_id,omitempty"`
}

var poolIsvVcAvailableRemindRequest = sync.Pool{
	New: func() any {
		return new(IsvVcAvailableRemindRequest)
	},
}

// GetIsvVcAvailableRemindRequest() 从对象池中获取IsvVcAvailableRemindRequest
func GetIsvVcAvailableRemindRequest() *IsvVcAvailableRemindRequest {
	return poolIsvVcAvailableRemindRequest.Get().(*IsvVcAvailableRemindRequest)
}

// ReleaseIsvVcAvailableRemindRequest 释放IsvVcAvailableRemindRequest
func ReleaseIsvVcAvailableRemindRequest(v *IsvVcAvailableRemindRequest) {
	v.VaccineId = ""
	v.PovId = ""
	poolIsvVcAvailableRemindRequest.Put(v)
}
