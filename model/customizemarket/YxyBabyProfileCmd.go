package customizemarket

import (
	"sync"
)

// YxyBabyProfileCmd 结构体
type YxyBabyProfileCmd struct {
	// 出生年月或者预产期
	Birthday string `json:"birthday,omitempty" xml:"birthday,omitempty"`
	// 性别
	Gender int64 `json:"gender,omitempty" xml:"gender,omitempty"`
	// 授权id
	ProfileId int64 `json:"profile_id,omitempty" xml:"profile_id,omitempty"`
	// 宝贝类型
	BabyType int64 `json:"baby_type,omitempty" xml:"baby_type,omitempty"`
}

var poolYxyBabyProfileCmd = sync.Pool{
	New: func() any {
		return new(YxyBabyProfileCmd)
	},
}

// GetYxyBabyProfileCmd() 从对象池中获取YxyBabyProfileCmd
func GetYxyBabyProfileCmd() *YxyBabyProfileCmd {
	return poolYxyBabyProfileCmd.Get().(*YxyBabyProfileCmd)
}

// ReleaseYxyBabyProfileCmd 释放YxyBabyProfileCmd
func ReleaseYxyBabyProfileCmd(v *YxyBabyProfileCmd) {
	v.Birthday = ""
	v.Gender = 0
	v.ProfileId = 0
	v.BabyType = 0
	poolYxyBabyProfileCmd.Put(v)
}
