package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机创建订单 APIResponse
taobao.bus.tvmcreateorder.set

提供给汽车票线下自助机的创建订单使用
*/
type TaobaoBusTvmcreateorderSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusTvmcreateorderSetResponse `json:"bus_tvmcreateorder_set_response,omitempty"` 
    TaobaoBusTvmcreateorderSetResponse
}

/* model for simplify = false
type TaobaoBusTvmcreateorderSetResponse struct {

    // alitripOrderId
    
    AlitripOrderId   string `json:"alitrip_order_id,omitempty"`
    

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusTvmcreateorderSetResponse struct {

    // alitripOrderId
    AlitripOrderId   string `json:"alitrip_order_id,omitempty"`

    // errorCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
