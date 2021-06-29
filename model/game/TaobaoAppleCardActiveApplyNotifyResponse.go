package game

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
苹果卡密申请激活回调接口 APIResponse
taobao.apple.card.active.apply.notify

苹果卡密申请激活回调接口
*/
type TaobaoAppleCardActiveApplyNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoAppleCardActiveApplyNotifyResponse
}

type TaobaoAppleCardActiveApplyNotifyResponse struct {
    XMLName xml.Name `xml:"apple_card_active_apply_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
    // 描述
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
}
