package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
线下自助机通知出票接口 APIResponse
taobao.bus.tvmbookorder.set

出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
*/
type TaobaoBusTvmbookorderSetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusTvmbookorderSetResponse `json:"bus_tvmbookorder_set_response,omitempty"` 
    TaobaoBusTvmbookorderSetResponse
}

/* model for simplify = false
type TaobaoBusTvmbookorderSetResponse struct {

    // errorCode
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // errorMsg
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusTvmbookorderSetResponse struct {

    // errorCode
    ResultCode   string `json:"result_code,omitempty"`

    // errorMsg
    ResultMsg   string `json:"result_msg,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

}
