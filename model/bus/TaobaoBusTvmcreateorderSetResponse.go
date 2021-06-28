package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机创建订单 APIResponse
taobao.bus.tvmcreateorder.set

提供给汽车票线下自助机的创建订单使用
*/
type TaobaoBusTvmcreateorderSetAPIResponse struct {
    model.CommonResponse
    TaobaoBusTvmcreateorderSetResponse
}

type TaobaoBusTvmcreateorderSetResponse struct {
    XMLName xml.Name `xml:"bus_tvmcreateorder_set_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // alitripOrderId
    
    AlitripOrderId   string `json:"alitrip_order_id,omitempty" xml:"alitrip_order_id,omitempty"`

    
    // errorCode
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
