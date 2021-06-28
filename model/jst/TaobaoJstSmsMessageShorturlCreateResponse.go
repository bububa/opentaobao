package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔营销效果短链生成 APIResponse
taobao.jst.sms.message.shorturl.create

聚石塔生成淘短链接口
*/
type TaobaoJstSmsMessageShorturlCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_sms_message_shorturl_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功
    
    RCode   string `json:"r_code,omitempty" xml:"