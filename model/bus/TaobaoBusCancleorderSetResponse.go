package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取消订单 APIResponse
taobao.bus.cancleorder.set

取消订单
*/
type TaobaoBusCancleorderSetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoBusCancleorderSetResponse `json:"taobao_bus_cancleorder_set_response,omitempty"`
}

type TaobaoBusCancleorderSetResponse struct {

    // 错误代码
    ErrorCode1   string `json:"error_code1,omitempty"`

    // 错误描述
    ErrorMsg1   string `json:"error_msg1,omitempty"`

    // success
    Success1   bool `json:"success1,omitempty"`

}
