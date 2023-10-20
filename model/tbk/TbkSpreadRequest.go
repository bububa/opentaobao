package tbk

import (
	"sync"
)

// TbkSpreadRequest 结构体
type TbkSpreadRequest struct {
	// 原始url, 只支持uland.taobao.com，s.click.taobao.com， ai.taobao.com，temai.taobao.com的域名转换，否则判错
	Url string `json:"url,omitempty" xml:"url,omitempty"`
}

var poolTbkSpreadRequest = sync.Pool{
	New: func() any {
		return new(TbkSpreadRequest)
	},
}

// GetTbkSpreadRequest() 从对象池中获取TbkSpreadRequest
func GetTbkSpreadRequest() *TbkSpreadRequest {
	return poolTbkSpreadRequest.Get().(*TbkSpreadRequest)
}

// ReleaseTbkSpreadRequest 释放TbkSpreadRequest
func ReleaseTbkSpreadRequest(v *TbkSpreadRequest) {
	v.Url = ""
	poolTbkSpreadRequest.Put(v)
}
