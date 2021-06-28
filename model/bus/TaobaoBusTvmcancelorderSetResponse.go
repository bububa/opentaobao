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
    TaobaoBusTvmcancelorderSetResponse
}

type TaobaoBusTvmcancelorderSetResponse struct {
    XMLName xml.Name `xml:"bus_tvmcancelorder_set_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码  ORDER_NOT_FOUND
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 错误描述  订单无法查询
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // true代表成功 false 代表失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
