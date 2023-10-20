package car

import (
	"sync"
)

// RentCarOrderDetailRsp 结构体
type RentCarOrderDetailRsp struct {
	// errorCode
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// errorMsg
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// buyerInfo
	BuyerInfo *BuyerInfo `json:"buyer_info,omitempty" xml:"buyer_info,omitempty"`
	// 信用免押信息
	DepositInfo *RentCarDepositInfo `json:"deposit_info,omitempty" xml:"deposit_info,omitempty"`
	// 订单信息
	OrderInfo *OrderDetailInfo `json:"order_info,omitempty" xml:"order_info,omitempty"`
	// 商家信息
	SellerInfo *SellerInfo `json:"seller_info,omitempty" xml:"seller_info,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolRentCarOrderDetailRsp = sync.Pool{
	New: func() any {
		return new(RentCarOrderDetailRsp)
	},
}

// GetRentCarOrderDetailRsp() 从对象池中获取RentCarOrderDetailRsp
func GetRentCarOrderDetailRsp() *RentCarOrderDetailRsp {
	return poolRentCarOrderDetailRsp.Get().(*RentCarOrderDetailRsp)
}

// ReleaseRentCarOrderDetailRsp 释放RentCarOrderDetailRsp
func ReleaseRentCarOrderDetailRsp(v *RentCarOrderDetailRsp) {
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.BuyerInfo = nil
	v.DepositInfo = nil
	v.OrderInfo = nil
	v.SellerInfo = nil
	v.Success = false
	poolRentCarOrderDetailRsp.Put(v)
}
