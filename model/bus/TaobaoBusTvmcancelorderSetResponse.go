package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机未付款取消订单 APIResponse
taobao.bus.tvmcancelorder.set

自助机汽车票未付款取消订单
*/
type TaobaoBusTvmcancelorderSetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"bus_tvmcancelorder_set_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码  ORDER_NOT_FOUND
    
    ResultCode   string `json:"result_code,omitempty" xml:"