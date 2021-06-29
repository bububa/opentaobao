package jst

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔短链信息查询 API返回值 
taobao.jst.sms.message.shorturl.query

聚石塔短链信息查询
*/
type TaobaoJstSmsMessageShorturlQueryAPIResponse struct {
    model.CommonResponse
    TaobaoJstSmsMessageShorturlQueryResponse
}

// 聚石塔短链信息查询 成功返回结果
type TaobaoJstSmsMessageShorturlQueryResponse struct {
    XMLName xml.Name `xml:"jst_sms_message_shorturl_query_response"`
    // 请求成功
    RCode   string `json:"r_code,omitempty" xml:"r_code,omitempty"`
    // 成功
    RSuccess   bool `json:"r_success,omitempty" xml:"r_success,omitempty"`
    // TOP请求ID
    RequestId   string `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 短链名
    ShortName   string `json:"short_name,omitempty" xml:"short_name,omitempty"`
    // 短链失效时间
    LifeEnd   string `json:"life_end,omitempty" xml:"life_end,omitempty"`
    // 短链生成时间
    LifeStart   string `json:"life_start,omitempty" xml:"life_start,omitempty"`
    // 查询短链信息失败
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
