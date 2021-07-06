package wdk

// MaochaoOrderQueryResult 结构体
type MaochaoOrderQueryResult struct {
	// 子订单列表
	SubOrderList []MaochaoWdkOrderDto `json:"sub_order_list,omitempty" xml:"sub_order_list>maochao_wdk_order_dto,omitempty"`
	// 返回码
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 返回码说明
	ReturnMsg string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
	// 是否调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}
