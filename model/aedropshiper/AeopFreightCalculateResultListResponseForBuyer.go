package aedropshiper

import (
	"sync"
)

// AeopFreightCalculateResultListResponseForBuyer 结构体
type AeopFreightCalculateResultListResponseForBuyer struct {
	// aeopFreightCalculateResultForBuyerDTOList
	AeopFreightCalculateResultForBuyerDTOList []AeopFreightCalculateResultForBuyerDto `json:"aeop_freight_calculate_result_for_buyer_d_t_o_list,omitempty" xml:"aeop_freight_calculate_result_for_buyer_d_t_o_list>aeop_freight_calculate_result_for_buyer_dto,omitempty"`
	// errorDesc
	ErrorDesc string `json:"error_desc,omitempty" xml:"error_desc,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAeopFreightCalculateResultListResponseForBuyer = sync.Pool{
	New: func() any {
		return new(AeopFreightCalculateResultListResponseForBuyer)
	},
}

// GetAeopFreightCalculateResultListResponseForBuyer() 从对象池中获取AeopFreightCalculateResultListResponseForBuyer
func GetAeopFreightCalculateResultListResponseForBuyer() *AeopFreightCalculateResultListResponseForBuyer {
	return poolAeopFreightCalculateResultListResponseForBuyer.Get().(*AeopFreightCalculateResultListResponseForBuyer)
}

// ReleaseAeopFreightCalculateResultListResponseForBuyer 释放AeopFreightCalculateResultListResponseForBuyer
func ReleaseAeopFreightCalculateResultListResponseForBuyer(v *AeopFreightCalculateResultListResponseForBuyer) {
	v.AeopFreightCalculateResultForBuyerDTOList = v.AeopFreightCalculateResultForBuyerDTOList[:0]
	v.ErrorDesc = ""
	v.Success = false
	poolAeopFreightCalculateResultListResponseForBuyer.Put(v)
}
