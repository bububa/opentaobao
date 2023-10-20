package customizemarket

import (
	"sync"
)

// ProfileQuery 结构体
type ProfileQuery struct {
	// 宠物名称
	PetNick string `json:"pet_nick,omitempty" xml:"pet_nick,omitempty"`
	// 每页数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 用户id
	UserId int64 `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 第几页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

var poolProfileQuery = sync.Pool{
	New: func() any {
		return new(ProfileQuery)
	},
}

// GetProfileQuery() 从对象池中获取ProfileQuery
func GetProfileQuery() *ProfileQuery {
	return poolProfileQuery.Get().(*ProfileQuery)
}

// ReleaseProfileQuery 释放ProfileQuery
func ReleaseProfileQuery(v *ProfileQuery) {
	v.PetNick = ""
	v.PageSize = 0
	v.UserId = 0
	v.PageNum = 0
	poolProfileQuery.Put(v)
}
