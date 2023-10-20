package iot

import (
	"sync"
)

// TaobaoAilabAicloudTopMessageListModel 结构体
type TaobaoAilabAicloudTopMessageListModel struct {
	// source
	Source string `json:"source,omitempty" xml:"source,omitempty"`
	// type
	Type string `json:"type,omitempty" xml:"type,omitempty"`
	// content
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// gmtModified
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// status
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolTaobaoAilabAicloudTopMessageListModel = sync.Pool{
	New: func() any {
		return new(TaobaoAilabAicloudTopMessageListModel)
	},
}

// GetTaobaoAilabAicloudTopMessageListModel() 从对象池中获取TaobaoAilabAicloudTopMessageListModel
func GetTaobaoAilabAicloudTopMessageListModel() *TaobaoAilabAicloudTopMessageListModel {
	return poolTaobaoAilabAicloudTopMessageListModel.Get().(*TaobaoAilabAicloudTopMessageListModel)
}

// ReleaseTaobaoAilabAicloudTopMessageListModel 释放TaobaoAilabAicloudTopMessageListModel
func ReleaseTaobaoAilabAicloudTopMessageListModel(v *TaobaoAilabAicloudTopMessageListModel) {
	v.Source = ""
	v.Type = ""
	v.Content = ""
	v.Url = ""
	v.GmtCreate = ""
	v.GmtModified = ""
	v.Status = 0
	v.Id = 0
	poolTaobaoAilabAicloudTopMessageListModel.Put(v)
}
