package game

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新用户签约通知接口 APIResponse
taobao.apple.newuser.sign.notify

用户付款成功后，资和信主动通知签约结果
*/
type TaobaoAppleNewuserSignNotifyAPIResponse struct {
    model.CommonResponse
    TaobaoAppleNewuserSignNotifyResponse
}

type TaobaoAppleNewuserSignNotifyResponse struct {
    XMLName xml.Name `xml:"apple_newuser_sign_notify_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 处理结果说明
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`

    
    // 处理结果码
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
