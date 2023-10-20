package tmallnr

import (
	"sync"
)

// NrtCrmActivityPageCreateReq 结构体
type NrtCrmActivityPageCreateReq struct {
	// isvCode
	IsvCode string `json:"isv_code,omitempty" xml:"isv_code,omitempty"`
	// 缩略图url
	Thumbnail string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	// 页面地址
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// 备注说明
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 简介
	Synopsis string `json:"synopsis,omitempty" xml:"synopsis,omitempty"`
	// 标题
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 业务code
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 页面ID
	PageId int64 `json:"page_id,omitempty" xml:"page_id,omitempty"`
	// 创建人id
	Creator int64 `json:"creator,omitempty" xml:"creator,omitempty"`
	// 页面类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolNrtCrmActivityPageCreateReq = sync.Pool{
	New: func() any {
		return new(NrtCrmActivityPageCreateReq)
	},
}

// GetNrtCrmActivityPageCreateReq() 从对象池中获取NrtCrmActivityPageCreateReq
func GetNrtCrmActivityPageCreateReq() *NrtCrmActivityPageCreateReq {
	return poolNrtCrmActivityPageCreateReq.Get().(*NrtCrmActivityPageCreateReq)
}

// ReleaseNrtCrmActivityPageCreateReq 释放NrtCrmActivityPageCreateReq
func ReleaseNrtCrmActivityPageCreateReq(v *NrtCrmActivityPageCreateReq) {
	v.IsvCode = ""
	v.Thumbnail = ""
	v.Url = ""
	v.Description = ""
	v.Synopsis = ""
	v.Title = ""
	v.BizCode = ""
	v.PageId = 0
	v.Creator = 0
	v.Type = 0
	poolNrtCrmActivityPageCreateReq.Put(v)
}
