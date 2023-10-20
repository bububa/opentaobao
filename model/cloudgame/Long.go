package cloudgame

import (
	"sync"
)

// Long 结构体
type Long struct {
	// 玩家id
	UserId string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// 位置索引
	PlayerIndex string `json:"player_index,omitempty" xml:"player_index,omitempty"`
}

var poolLong = sync.Pool{
	New: func() any {
		return new(Long)
	},
}

// GetLong() 从对象池中获取Long
func GetLong() *Long {
	return poolLong.Get().(*Long)
}

// ReleaseLong 释放Long
func ReleaseLong(v *Long) {
	v.UserId = ""
	v.PlayerIndex = ""
	poolLong.Put(v)
}
