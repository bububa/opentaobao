package game

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密取消激活回调接口 API返回值 
taobao.apple.card.active.cancel.notify

苹果卡密取消激活回调接口
*/
type TaobaoAppleCardActiveCancelNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoAppleCardActiveCancelNotifyResponse
}

// 苹果卡密取消激活回调接口 成功返回结果
type TaobaoAppleCardActiveCancelNotifyResponse struct {
    XMLName xml.Name `xml:"apple_card_active_cancel_notify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 描述
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
