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
	RequestId     string         `json:"request_id,omitempty" xml:"bus_historyorder_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"