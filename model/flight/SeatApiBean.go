package flight

import (
	"sync"
)

// SeatApiBean 结构体
type SeatApiBean struct {
	// 座位列 1:靠过道，2:靠窗，3:并排
	SeatArea int64 `json:"seat_area,omitempty" xml:"seat_area,omitempty"`
	// 排次 1:前排，2:中排，3:后排
	SeatRow int64 `json:"seat_row,omitempty" xml:"seat_row,omitempty"`
}

var poolSeatApiBean = sync.Pool{
	New: func() any {
		return new(SeatApiBean)
	},
}

// GetSeatApiBean() 从对象池中获取SeatApiBean
func GetSeatApiBean() *SeatApiBean {
	return poolSeatApiBean.Get().(*SeatApiBean)
}

// ReleaseSeatApiBean 释放SeatApiBean
func ReleaseSeatApiBean(v *SeatApiBean) {
	v.SeatArea = 0
	v.SeatRow = 0
	poolSeatApiBean.Put(v)
}
