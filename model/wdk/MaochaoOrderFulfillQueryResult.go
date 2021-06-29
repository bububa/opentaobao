package wdk

// MaochaoOrderFulfillQueryResult 
type MaochaoOrderFulfillQueryResult struct {
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
    // 返回码
    ReturnCode   string `json:"return_code,omitempty" xml:"return_code,omitempty"`
    // 返回码说明
    ReturnMsg   string `json:"return_msg,omitempty" xml:"return_msg,omitempty"`
    // 履约单列表
    FulfillOrderList   []MaochaoWdkOrderFulfillDTO `json:"fulfill_order_list,omitempty" xml:"fulfill_order_list>maochao_wdk_order_fulfill_dto,omitempty"`
}
