package jst

import (
    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短链信息查询 APIResponse
taobao.jst.sms.message.shorturl.query

聚石塔短链信息查询
*/
type TaobaoJstSmsMessageShorturlQueryAPIResponse struct {
    model.CommonResponse
    Response *TaobaoJstSmsMessageShorturlQueryResponse `json:"taobao_jst_sms_message_shorturl_query_response,omitempty"`
}

type TaobaoJstSmsMessageShorturlQueryResponse struct {

    // 请求成功
    RCode   string `json:"r_code,omitempty"`

    // 成功
    RSuccess   bool `json:"r_success,omitempty"`

    // TOP请求ID
    RequestId   string `json:"request_id,omitempty"`

    // 短链名
    ShortName   string `json:"short_name,omitempty"`

    // 短链失效时间
    LifeEnd   string `json:"life_end,omitempty"`

    // 短链生成时间
    LifeStart   string `json:"life_start,omitempty"`

    // 查询短链信息失败
    Message   string `json:"message,omitempty"`

}
