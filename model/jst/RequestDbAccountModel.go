package jst

import (
	"sync"
)

// RequestDbAccountModel 结构体
type RequestDbAccountModel struct {
	// 账户描述
	AccountDesc string `json:"account_desc,omitempty" xml:"account_desc,omitempty"`
	// 账户名称
	AccountName string `json:"account_name,omitempty" xml:"account_name,omitempty"`
	// 新建db名称
	DbName string `json:"db_name,omitempty" xml:"db_name,omitempty"`
	// rds实例名称
	InstanceName string `json:"instance_name,omitempty" xml:"instance_name,omitempty"`
	// 账户密码
	Password string `json:"password,omitempty" xml:"password,omitempty"`
}

var poolRequestDbAccountModel = sync.Pool{
	New: func() any {
		return new(RequestDbAccountModel)
	},
}

// GetRequestDbAccountModel() 从对象池中获取RequestDbAccountModel
func GetRequestDbAccountModel() *RequestDbAccountModel {
	return poolRequestDbAccountModel.Get().(*RequestDbAccountModel)
}

// ReleaseRequestDbAccountModel 释放RequestDbAccountModel
func ReleaseRequestDbAccountModel(v *RequestDbAccountModel) {
	v.AccountDesc = ""
	v.AccountName = ""
	v.DbName = ""
	v.InstanceName = ""
	v.Password = ""
	poolRequestDbAccountModel.Put(v)
}
