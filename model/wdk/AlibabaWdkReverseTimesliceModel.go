package wdk

import (
	"sync"
)

// AlibabaWdkReverseTimesliceModel 结构体
type AlibabaWdkReverseTimesliceModel struct {
	// 时间片对象
	DateTimeSliceCollectionDTOList []DateTimeSliceCollectionDtoList `json:"date_time_slice_collection_d_t_o_list,omitempty" xml:"date_time_slice_collection_d_t_o_list>date_time_slice_collection_dto_list,omitempty"`
	// 子单号
	SubOutOrderId string `json:"sub_out_order_id,omitempty" xml:"sub_out_order_id,omitempty"`
}

var poolAlibabaWdkReverseTimesliceModel = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseTimesliceModel)
	},
}

// GetAlibabaWdkReverseTimesliceModel() 从对象池中获取AlibabaWdkReverseTimesliceModel
func GetAlibabaWdkReverseTimesliceModel() *AlibabaWdkReverseTimesliceModel {
	return poolAlibabaWdkReverseTimesliceModel.Get().(*AlibabaWdkReverseTimesliceModel)
}

// ReleaseAlibabaWdkReverseTimesliceModel 释放AlibabaWdkReverseTimesliceModel
func ReleaseAlibabaWdkReverseTimesliceModel(v *AlibabaWdkReverseTimesliceModel) {
	v.DateTimeSliceCollectionDTOList = v.DateTimeSliceCollectionDTOList[:0]
	v.SubOutOrderId = ""
	poolAlibabaWdkReverseTimesliceModel.Put(v)
}
