package game

import (
    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密申请激活回调接口 APIResponse
taobao.apple.card.active.apply.notify

苹果卡密申请激活回调接口
*/
type TaobaoAppleCardActiveApplyNotifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAppleCardActiveApplyNotifyResponse `json:"taobao_apple_card_active_apply_notify_response,omitempty"`
}

type TaobaoAppleCardActiveApplyNotifyResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 描述
    ResultMsg   string `json:"result_msg,omitempty"`

}
