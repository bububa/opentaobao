package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家汽车票车次更新通知接口 APIResponse
taobao.bus.busnumber.set

商家汽车票车次更新后，调用该接口通知平台。
*/
type TaobaoBusBusnumberSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusBusnumberSetResponse `json:"bus_busnumber_set_response,omitempty"` 
    TaobaoBusBusnumberSetResponse
}

/* model for simplify = false
type TaobaoBusBusnumberSetResponse struct {

    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 返回对象
    
    Module   string `json:"module,omitempty"`
    

    // 错误编码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 接口调用是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusBusnumberSetResponse struct {

    // 错误描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 返回对象
    Module   string `json:"module,omitempty"`

    // 错误编码
    ResultCode   string `json:"result_code,omitempty"`

    // 接口调用是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
