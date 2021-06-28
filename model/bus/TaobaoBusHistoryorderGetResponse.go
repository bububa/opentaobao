package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
历史订单查询（对账） APIResponse
taobao.bus.historyorder.get

历史订单查询，对账接口
*/
type TaobaoBusHistoryorderGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusHistoryorderGetResponse
}

type TaobaoBusHistoryorderGetResponse struct {
    XMLName xml.Name `xml:"bus_historyorder_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // errorMsg 错误原因
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // busoMainOrderHistoryPageVO 订单详情
    
    BusoMainOrderHistoryPageVO   *BusoMainOrderHistoryPageVo `json:"buso_main_order_history_page_v_o,omitempty" xml:"buso_main_order_history_page_v_o,omitempty"`

    
    // success true 成功  false失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
