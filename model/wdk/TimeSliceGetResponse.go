package wdk

// TimeSliceGetResponse 结构体
type TimeSliceGetResponse struct {
	// 时间片列表
	TimeSliceList []PromiseTimeSlice `json:"time_slice_list,omitempty" xml:"time_slice_list>promise_time_slice,omitempty"`
	// 商品信息
	ProductList []FulfillProduct `json:"product_list,omitempty" xml:"product_list>fulfill_product,omitempty"`
}
