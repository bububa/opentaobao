package game

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新用户签约结果通知接口v2 APIResponse
taobao.apple.newuser.sign.notify.newversion

资和信主动通知签约结果
*/
type TaobaoAppleNewuserSignNotifyNewversionAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"apple_newuser_sign_notify_newversion_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 处理结果说明
    
    ResultMsg   string `json:"result_msg,omitempty" xml:"