package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
自助机失效二维码 APIResponse
taobao.bus.disableqrcode.set

使创建的二维码失效
*/
type TaobaoBusDisableqrcodeSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusDisableqrcodeSetResponse `json:"bus_disableqrcode_set_response,omitempty"` 
    TaobaoBusDisableqrcodeSetResponse
}

/* model for simplify = false
type TaobaoBusDisableqrcodeSetResponse struct {

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusDisableqrcodeSetResponse struct {

    // errorCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
