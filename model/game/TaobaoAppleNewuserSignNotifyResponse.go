package game

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新用户签约通知接口 APIResponse
taobao.apple.newuser.sign.notify

用户付款成功后，资和信主动通知签约结果
*/
type TaobaoAppleNewuserSignNotifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAppleNewuserSignNotifyResponse `json:"taobao_apple_newuser_sign_notify_response,omitempty"`
}

type TaobaoAppleNewuserSignNotifyResponse struct {

    // 处理结果说明
    ResultMsg   string `json:"result_msg,omitempty"`

    // 处理结果码
    ResultCode   string `json:"result_code,omitempty"`

}
