package aedropshiper

// PlaceOrderRes4OpenApiDto 结构体
type PlaceOrderRes4OpenApiDto struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// orderList
	OrderList []int64 `json:"order_list,omitempty" xml:"order_list>int64,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
