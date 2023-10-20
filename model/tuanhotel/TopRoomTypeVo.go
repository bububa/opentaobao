package tuanhotel

import (
	"sync"
)

// TopRoomTypeVo 结构体
type TopRoomTypeVo struct {
	// 房型名称
	TypeName string `json:"type_name,omitempty" xml:"type_name,omitempty"`
	// 房型ID
	TypeId int64 `json:"type_id,omitempty" xml:"type_id,omitempty"`
}

var poolTopRoomTypeVo = sync.Pool{
	New: func() any {
		return new(TopRoomTypeVo)
	},
}

// GetTopRoomTypeVo() 从对象池中获取TopRoomTypeVo
func GetTopRoomTypeVo() *TopRoomTypeVo {
	return poolTopRoomTypeVo.Get().(*TopRoomTypeVo)
}

// ReleaseTopRoomTypeVo 释放TopRoomTypeVo
func ReleaseTopRoomTypeVo(v *TopRoomTypeVo) {
	v.TypeName = ""
	v.TypeId = 0
	poolTopRoomTypeVo.Put(v)
}
