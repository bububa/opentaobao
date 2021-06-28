package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机查询订单信息 APIResponse
taobao.bus.tvmqueryorder.get

查询订单详情
*/
type TaobaoBusTvmqueryorderGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusTvmqueryorderGetResponse `json:"bus_tvmqueryorder_get_response,omitempty"` 
    TaobaoBusTvmqueryorderGetResponse
}

/* model for simplify = false
type TaobaoBusTvmqueryorderGetResponse struct {

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // tvmBusOrderLineInfo
    
    TvmBusOrderLineInfo  *struct {
        TvmBusOrderLineInfo  *TvmBusOrderLineInfo `json:"tvm_bus_order_line_info,omitempty"`
    } `json:"tvm_bus_order_line_info,omitempty"`
    

}
*/

type TaobaoBusTvmqueryorderGetResponse struct {

    // errorCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // tvmBusOrderLineInfo
    TvmBusOrderLineInfo   *TvmBusOrderLineInfo `json:"tvm_bus_order_line_info,omitempty"`

}
