package tbk

import (
	"sync"
)

// EventUrlList 结构体
type EventUrlList struct {
	// 物料对应错误描述
	Msg string `json:"msg,omitempty" xml:"msg,omitempty"`
	// 入参的会场ID
	InputPageId string `json:"input_page_id,omitempty" xml:"input_page_id,omitempty"`
	// 物料对应错误码
	Code int64 `json:"code,omitempty" xml:"code,omitempty"`
	// 链接信息
	LinkInfoDto *LinkInfoDto `json:"link_info_dto,omitempty" xml:"link_info_dto,omitempty"`
}

var poolEventUrlList = sync.Pool{
	New: func() any {
		return new(EventUrlList)
	},
}

// GetEventUrlList() 从对象池中获取EventUrlList
func GetEventUrlList() *EventUrlList {
	return poolEventUrlList.Get().(*EventUrlList)
}

// ReleaseEventUrlList 释放EventUrlList
func ReleaseEventUrlList(v *EventUrlList) {
	v.Msg = ""
	v.InputPageId = ""
	v.Code = 0
	v.LinkInfoDto = nil
	poolEventUrlList.Put(v)
}
