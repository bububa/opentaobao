package game

import (
    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密取消激活回调接口 APIResponse
taobao.apple.card.active.cancel.notify

苹果卡密取消激活回调接口
*/
type TaobaoAppleCardActiveCancelNotifyAPIResponse struct {
    model.CommonResponse
    Response *TaobaoAppleCardActiveCancelNotifyResponse `json:"taobao_apple_card_active_cancel_notify_response,omitempty"`
}

type TaobaoAppleCardActiveCancelNotifyResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 描述
    ResultMsg   string `json:"result_msg,omitempty"`

}
