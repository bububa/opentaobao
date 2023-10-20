package tmallcms

import (
	"sync"
)

// SpreadLinkDo 结构体
type SpreadLinkDo struct {
	// 修改时间
	GmtModified string `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	// 推广文案
	Adword string `json:"adword,omitempty" xml:"adword,omitempty"`
	// 推广信息
	Adinfo string `json:"adinfo,omitempty" xml:"adinfo,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 微信推广链接
	Adwxurl string `json:"adwxurl,omitempty" xml:"adwxurl,omitempty"`
	// 二维码地址
	Qrcode string `json:"qrcode,omitempty" xml:"qrcode,omitempty"`
	// 推广链接
	Adurl string `json:"adurl,omitempty" xml:"adurl,omitempty"`
	// 源链接
	Url string `json:"url,omitempty" xml:"url,omitempty"`
	// ID
	Id int64 `json:"id,omitempty" xml:"id,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 商家ID
	SellerId int64 `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 类型
	Type int64 `json:"type,omitempty" xml:"type,omitempty"`
}

var poolSpreadLinkDo = sync.Pool{
	New: func() any {
		return new(SpreadLinkDo)
	},
}

// GetSpreadLinkDo() 从对象池中获取SpreadLinkDo
func GetSpreadLinkDo() *SpreadLinkDo {
	return poolSpreadLinkDo.Get().(*SpreadLinkDo)
}

// ReleaseSpreadLinkDo 释放SpreadLinkDo
func ReleaseSpreadLinkDo(v *SpreadLinkDo) {
	v.GmtModified = ""
	v.Adword = ""
	v.Adinfo = ""
	v.GmtCreate = ""
	v.Adwxurl = ""
	v.Qrcode = ""
	v.Adurl = ""
	v.Url = ""
	v.Id = 0
	v.Status = 0
	v.SellerId = 0
	v.Type = 0
	poolSpreadLinkDo.Put(v)
}
