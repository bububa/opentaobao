package bus

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家汽车票车次更新通知接口 APIResponse
taobao.bus.busnumber.set

商家汽车票车次更新后，调用该接口通知平台。
*/
type TaobaoBusBusnumberSetAPIResponse struct {
    model.CommonResponse
    TaobaoBusBusnumberSetResponse
}

type TaobaoBusBusnumberSetResponse struct {
    XMLName xml.Name `xml:"bus_busnumber_set_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 返回对象
    
    Module   string `json:"module,omitempty" xml:"module,omitempty"`

    
    // 错误编码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 接口调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
