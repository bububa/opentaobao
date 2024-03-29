package xhotelonlineorder

import (
	"sync"
)

// UpdateRateDo 结构体
type UpdateRateDo struct {
	// 入离日期
	CheckDateDOList []CheckDateDo `json:"check_date_d_o_list,omitempty" xml:"check_date_d_o_list>check_date_do,omitempty"`
	// 酒店id
	OutHid string `json:"out_hid,omitempty" xml:"out_hid,omitempty"`
	// 儿童年龄，多个儿童年龄用竖线分割：4|5，默认儿童年龄为5岁
	Ages string `json:"ages,omitempty" xml:"ages,omitempty"`
	// 成人数
	Adults int64 `json:"adults,omitempty" xml:"adults,omitempty"`
	// 儿童数
	Children int64 `json:"children,omitempty" xml:"children,omitempty"`
}

var poolUpdateRateDo = sync.Pool{
	New: func() any {
		return new(UpdateRateDo)
	},
}

// GetUpdateRateDo() 从对象池中获取UpdateRateDo
func GetUpdateRateDo() *UpdateRateDo {
	return poolUpdateRateDo.Get().(*UpdateRateDo)
}

// ReleaseUpdateRateDo 释放UpdateRateDo
func ReleaseUpdateRateDo(v *UpdateRateDo) {
	v.CheckDateDOList = v.CheckDateDOList[:0]
	v.OutHid = ""
	v.Ages = ""
	v.Adults = 0
	v.Children = 0
	poolUpdateRateDo.Put(v)
}
