package tmallsc

import (
	"sync"
)

// UpdateServiceProgressRequest 结构体
type UpdateServiceProgressRequest struct {
	// 图片地址回传集合
	PicUrlList []string `json:"pic_url_list,omitempty" xml:"pic_url_list>string,omitempty"`
	// 服务描述
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 服务动作
	Action string `json:"action,omitempty" xml:"action,omitempty"`
	// 扩展信息回传备注
	AttributeMap string `json:"attribute_map,omitempty" xml:"attribute_map,omitempty"`
	// 工单id
	WorkcardId int64 `json:"workcard_id,omitempty" xml:"workcard_id,omitempty"`
}

var poolUpdateServiceProgressRequest = sync.Pool{
	New: func() any {
		return new(UpdateServiceProgressRequest)
	},
}

// GetUpdateServiceProgressRequest() 从对象池中获取UpdateServiceProgressRequest
func GetUpdateServiceProgressRequest() *UpdateServiceProgressRequest {
	return poolUpdateServiceProgressRequest.Get().(*UpdateServiceProgressRequest)
}

// ReleaseUpdateServiceProgressRequest 释放UpdateServiceProgressRequest
func ReleaseUpdateServiceProgressRequest(v *UpdateServiceProgressRequest) {
	v.PicUrlList = v.PicUrlList[:0]
	v.Desc = ""
	v.Action = ""
	v.AttributeMap = ""
	v.WorkcardId = 0
	poolUpdateServiceProgressRequest.Put(v)
}
