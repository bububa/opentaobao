package wdk

// AlibabaWdkReverseTimesliceModel 结构体
type AlibabaWdkReverseTimesliceModel struct {
	// 时间片对象
	DateTimeSliceCollectionDTOList []DateTimeSliceCollectionDtoList `json:"date_time_slice_collection_d_t_o_list,omitempty" xml:"date_time_slice_collection_d_t_o_list>date_time_slice_collection_dto_list,omitempty"`
	// 子单号
	SubOutOrderId string `json:"sub_out_order_id,omitempty" xml:"sub_out_order_id,omitempty"`
}
