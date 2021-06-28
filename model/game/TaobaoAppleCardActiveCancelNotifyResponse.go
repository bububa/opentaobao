package game

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密取消激活回调接口 APIResponse
taobao.apple.card.active.cancel.notify

苹果卡密取消激活回调接口
*/
type TaobaoAppleCardActiveCancelNotifyAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"apple_card_active_cancel_notify_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"