package tvupadmin

import (
	"sync"
)

// AccountSuggestDo 结构体
type AccountSuggestDo struct {
	// 用户昵称
	Nick string `json:"nick,omitempty" xml:"nick,omitempty"`
	// 优酷账号昵称
	YtNick string `json:"yt_nick,omitempty" xml:"yt_nick,omitempty"`
	// 用户ID
	Uid int64 `json:"uid,omitempty" xml:"uid,omitempty"`
	// 优酷账号ID
	Ytid int64 `json:"ytid,omitempty" xml:"ytid,omitempty"`
}

var poolAccountSuggestDo = sync.Pool{
	New: func() any {
		return new(AccountSuggestDo)
	},
}

// GetAccountSuggestDo() 从对象池中获取AccountSuggestDo
func GetAccountSuggestDo() *AccountSuggestDo {
	return poolAccountSuggestDo.Get().(*AccountSuggestDo)
}

// ReleaseAccountSuggestDo 释放AccountSuggestDo
func ReleaseAccountSuggestDo(v *AccountSuggestDo) {
	v.Nick = ""
	v.YtNick = ""
	v.Uid = 0
	v.Ytid = 0
	poolAccountSuggestDo.Put(v)
}
