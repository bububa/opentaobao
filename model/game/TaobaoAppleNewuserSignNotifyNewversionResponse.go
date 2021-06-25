package game

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新用户签约结果通知接口v2 APIResponse
taobao.apple.newuser.sign.notify.newversion

资和信主动通知签约结果
*/
type TaobaoAppleNewuserSignNotifyNewversionAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAppleNewuserSignNotifyNewversionResponse `json:"taobao_apple_newuser_sign_notify_newversion_response,omitempty"`
}

type TaobaoAppleNewuserSignNotifyNewversionResponse struct {

    // 处理结果说明
    ResultMsg   string `json:"result_msg,omitempty"`

    // 处理结果码
    ResultCode   string `json:"result_code,omitempty"`

}
