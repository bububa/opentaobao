package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
历史订单查询（对账） APIResponse
taobao.bus.historyorder.get

历史订单查询，对账接口
*/
type TaobaoBusHistoryorderGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusHistoryorderGetResponse `json:"bus_historyorder_get_response,omitempty"` 
    TaobaoBusHistoryorderGetResponse
}

/* model for simplify = false
type TaobaoBusHistoryorderGetResponse struct {

    // errorCode 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg 错误原因
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // busoMainOrderHistoryPageVO 订单详情
    
    BusoMainOrderHistoryPageVO  *struct {
        BusoMainOrderHistoryPageVo  *BusoMainOrderHistoryPageVo `json:"buso_main_order_history_page_vo,omitempty"`
    } `json:"buso_main_order_history_page_v_o,omitempty"`
    

    // success true 成功  false失败
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusHistoryorderGetResponse struct {

    // errorCode 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg 错误原因
    ResultMsg   string `json:"result_msg,omitempty"`

    // busoMainOrderHistoryPageVO 订单详情
    BusoMainOrderHistoryPageVO   *BusoMainOrderHistoryPageVo `json:"buso_main_order_history_page_v_o,omitempty"`

    // success true 成功  false失败
    IsSuccess   bool `json:"is_success,omitempty"`

}
