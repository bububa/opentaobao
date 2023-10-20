package perfect

import (
	"sync"
)

// BoxCodeRequest 结构体
type BoxCodeRequest struct {
	// 箱号类型
	BoxCodeType string `json:"box_code_type,omitempty" xml:"box_code_type,omitempty"`
	// 容器号
	ContainerCode string `json:"container_code,omitempty" xml:"container_code,omitempty"`
	// 业务类型
	BizType string `json:"biz_type,omitempty" xml:"biz_type,omitempty"`
}

var poolBoxCodeRequest = sync.Pool{
	New: func() any {
		return new(BoxCodeRequest)
	},
}

// GetBoxCodeRequest() 从对象池中获取BoxCodeRequest
func GetBoxCodeRequest() *BoxCodeRequest {
	return poolBoxCodeRequest.Get().(*BoxCodeRequest)
}

// ReleaseBoxCodeRequest 释放BoxCodeRequest
func ReleaseBoxCodeRequest(v *BoxCodeRequest) {
	v.BoxCodeType = ""
	v.ContainerCode = ""
	v.BizType = ""
	poolBoxCodeRequest.Put(v)
}
