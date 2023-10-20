package traderate

import (
	"sync"
)

// MixRateVo 结构体
type MixRateVo struct {
	// 图片信息
	PictureUrls []string `json:"picture_urls,omitempty" xml:"picture_urls>string,omitempty"`
	// 正文内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
	// 创建时间
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// 正文摘要
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 用户头像
	UserIcon string `json:"user_icon,omitempty" xml:"user_icon,omitempty"`
	// 用户昵称
	UserNick string `json:"user_nick,omitempty" xml:"user_nick,omitempty"`
	// 总评分
	TotalScore int64 `json:"total_score,omitempty" xml:"total_score,omitempty"`
}

var poolMixRateVo = sync.Pool{
	New: func() any {
		return new(MixRateVo)
	},
}

// GetMixRateVo() 从对象池中获取MixRateVo
func GetMixRateVo() *MixRateVo {
	return poolMixRateVo.Get().(*MixRateVo)
}

// ReleaseMixRateVo 释放MixRateVo
func ReleaseMixRateVo(v *MixRateVo) {
	v.PictureUrls = v.PictureUrls[:0]
	v.Content = ""
	v.GmtCreate = ""
	v.Title = ""
	v.UserIcon = ""
	v.UserNick = ""
	v.TotalScore = 0
	poolMixRateVo.Put(v)
}
