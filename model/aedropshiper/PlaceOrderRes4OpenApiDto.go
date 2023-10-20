package aedropshiper

import (
	"sync"
)

// PlaceOrderRes4OpenApiDto 结构体
type PlaceOrderRes4OpenApiDto struct {
	// orderList
	OrderList []int64 `json:"order_list,omitempty" xml:"order_list>int64,omitempty"`
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

var poolPlaceOrderRes4OpenApiDto = sync.Pool{
	New: func() any {
		return new(PlaceOrderRes4OpenApiDto)
	},
}

// GetPlaceOrderRes4OpenApiDto() 从对象池中获取PlaceOrderRes4OpenApiDto
func GetPlaceOrderRes4OpenApiDto() *PlaceOrderRes4OpenApiDto {
	return poolPlaceOrderRes4OpenApiDto.Get().(*PlaceOrderRes4OpenApiDto)
}

// ReleasePlaceOrderRes4OpenApiDto 释放PlaceOrderRes4OpenApiDto
func ReleasePlaceOrderRes4OpenApiDto(v *PlaceOrderRes4OpenApiDto) {
	v.OrderList = v.OrderList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.IsSuccess = false
	poolPlaceOrderRes4OpenApiDto.Put(v)
}
