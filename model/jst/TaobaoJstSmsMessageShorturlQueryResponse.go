package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短链信息查询 APIResponse
taobao.jst.sms.message.shorturl.query

聚石塔短链信息查询
*/
type TaobaoJstSmsMessageShorturlQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jst_sms_message_shorturl_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 请求成功
    
    RCode   string `json:"r_code,omitempty" xml:"