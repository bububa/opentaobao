package qimen

import (
	"sync"
)

// TaobaoQimenOrderQueryRequest 结构体
type TaobaoQimenOrderQueryRequest struct {
	// 姓名, string (50) , 必填
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	// 开始时间, string (19) , YYYY-MM-DD HH:MM:SS
	StartTime string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// 结束时间, string (19) , YYYY-MM-DD HH:MM:SS
	EndTime string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 固定电话, string (50)
	Tel string `json:"tel,omitempty" xml:"tel,omitempty"`
	// 移动电话, string (50) , 必填
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 1001，客服咨询；1002，售后服务
	Scene string `json:"scene,omitempty" xml:"scene,omitempty"`
}

var poolTaobaoQimenOrderQueryRequest = sync.Pool{
	New: func() any {
		return new(TaobaoQimenOrderQueryRequest)
	},
}

// GetTaobaoQimenOrderQueryRequest() 从对象池中获取TaobaoQimenOrderQueryRequest
func GetTaobaoQimenOrderQueryRequest() *TaobaoQimenOrderQueryRequest {
	return poolTaobaoQimenOrderQueryRequest.Get().(*TaobaoQimenOrderQueryRequest)
}

// ReleaseTaobaoQimenOrderQueryRequest 释放TaobaoQimenOrderQueryRequest
func ReleaseTaobaoQimenOrderQueryRequest(v *TaobaoQimenOrderQueryRequest) {
	v.Name = ""
	v.StartTime = ""
	v.EndTime = ""
	v.Tel = ""
	v.Mobile = ""
	v.Scene = ""
	poolTaobaoQimenOrderQueryRequest.Put(v)
}
