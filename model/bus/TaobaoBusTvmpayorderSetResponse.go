package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
自助机条形码被动支付 APIResponse
taobao.bus.tvmpayorder.set

汽车票线下自助机条形码支付
*/
type TaobaoBusTvmpayorderSetAPIResponse struct {
    model.CommonResponse
    TaobaoBusTvmpayorderSetResponse
}

type TaobaoBusTvmpayorderSetResponse struct {
    XMLName xml.Name `xml:"bus_tvmpayorder_set_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // errorCode  线下扫码支付  错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // errorMsg 错误信息
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // success true 成功 false 失败
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // payTime
    
    PayTime   string `json:"pay_time,omitempty" xml:"pay_time,omitempty"`

    
}
