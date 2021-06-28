package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机查询订单信息 APIResponse
taobao.bus.tvmqueryorder.get

查询订单详情
*/
type TaobaoBusTvmqueryorderGetAPIResponse struct {
    model.CommonResponse
    TaobaoBusTvmqueryorderGetResponse
}

type TaobaoBusTvmqueryorderGetResponse struct {
    XMLName xml.Name `xml:"bus_tvmqueryorder_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // tvmBusOrderLineInfo
    
    TvmBusOrderLineInfo   *TvmBusOrderLineInfo `json:"tvm_bus_order_line_info,omitempty" xml:"tvm_bus_order_line_info,omitempty"`

    
}
