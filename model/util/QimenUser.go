package util

import (
	"sync"
)

// QimenUser 结构体
type QimenUser struct {
	// memo
	Memo string `json:"memo,omitempty" xml:"memo,omitempty"`
	// gmtCreate
	GmtCreate string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// sellerNick
	SellerNick string `json:"seller_nick,omitempty" xml:"seller_nick,omitempty"`
}

var poolQimenUser = sync.Pool{
	New: func() any {
		return new(QimenUser)
	},
}

// GetQimenUser() 从对象池中获取QimenUser
func GetQimenUser() *QimenUser {
	return poolQimenUser.Get().(*QimenUser)
}

// ReleaseQimenUser 释放QimenUser
func ReleaseQimenUser(v *QimenUser) {
	v.Memo = ""
	v.GmtCreate = ""
	v.SellerNick = ""
	poolQimenUser.Put(v)
}
