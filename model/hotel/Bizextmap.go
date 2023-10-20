package hotel

import (
	"sync"
)

// Bizextmap 结构体
type Bizextmap struct {
	// empty
	Empty bool `json:"empty,omitempty" xml:"empty,omitempty"`
}

var poolBizextmap = sync.Pool{
	New: func() any {
		return new(Bizextmap)
	},
}

// GetBizextmap() 从对象池中获取Bizextmap
func GetBizextmap() *Bizextmap {
	return poolBizextmap.Get().(*Bizextmap)
}

// ReleaseBizextmap 释放Bizextmap
func ReleaseBizextmap(v *Bizextmap) {
	v.Empty = false
	poolBizextmap.Put(v)
}
