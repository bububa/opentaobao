package product

import (
	"sync"
)

// DapeiDo 结构体
type DapeiDo struct {
	// items
	Items []DapeiTemplateItem `json:"items,omitempty" xml:"items>dapei_template_item,omitempty"`
	// title
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// desc
	Desc string `json:"desc,omitempty" xml:"desc,omitempty"`
	// url
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// id
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
}

var poolDapeiDo = sync.Pool{
	New: func() any {
		return new(DapeiDo)
	},
}

// GetDapeiDo() 从对象池中获取DapeiDo
func GetDapeiDo() *DapeiDo {
	return poolDapeiDo.Get().(*DapeiDo)
}

// ReleaseDapeiDo 释放DapeiDo
func ReleaseDapeiDo(v *DapeiDo) {
	v.Items = v.Items[:0]
	v.Title = ""
	v.Desc = ""
	v.Url = ""
	v.Id = 0
	poolDapeiDo.Put(v)
}
