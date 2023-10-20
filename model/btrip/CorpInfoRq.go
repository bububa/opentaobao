package btrip

import (
	"sync"
)

// CorpInfoRq 结构体
type CorpInfoRq struct {
	// 第三方企业ID（注册签约时必填）
	CorpId string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	// 企业名称（注册签约时必填）
	CorpName string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
	// 行业
	Industry string `json:"industry,omitempty" xml:"industry,omitempty"`
	// 省
	Province string `json:"province,omitempty" xml:"province,omitempty"`
	// 市
	City string `json:"city,omitempty" xml:"city,omitempty"`
	// 企业联系人
	Contact string `json:"contact,omitempty" xml:"contact,omitempty"`
	// 企业联系电话
	Mobile string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 差旅规模：1代表5万以下，2代表5-10万，3代表10-50万，4代表50万以上
	Scope int64 `json:"scope,omitempty" xml:"scope,omitempty"`
	// 企业人数
	PeopleSize int64 `json:"people_size,omitempty" xml:"people_size,omitempty"`
}

var poolCorpInfoRq = sync.Pool{
	New: func() any {
		return new(CorpInfoRq)
	},
}

// GetCorpInfoRq() 从对象池中获取CorpInfoRq
func GetCorpInfoRq() *CorpInfoRq {
	return poolCorpInfoRq.Get().(*CorpInfoRq)
}

// ReleaseCorpInfoRq 释放CorpInfoRq
func ReleaseCorpInfoRq(v *CorpInfoRq) {
	v.CorpId = ""
	v.CorpName = ""
	v.Industry = ""
	v.Province = ""
	v.City = ""
	v.Contact = ""
	v.Mobile = ""
	v.Scope = 0
	v.PeopleSize = 0
	poolCorpInfoRq.Put(v)
}
