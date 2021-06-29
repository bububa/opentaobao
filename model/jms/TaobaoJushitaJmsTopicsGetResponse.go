package jms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户nick获取开通的消息列表 API返回值 
taobao.jushita.jms.topics.get

根据用户nick获取开通的消息列表
*/
type TaobaoJushitaJmsTopicsGetAPIResponse struct {
    model.CommonResponse
    TaobaoJushitaJmsTopicsGetResponse
}

// 根据用户nick获取开通的消息列表 成功返回结果
type TaobaoJushitaJmsTopicsGetResponse struct {
    XMLName xml.Name `xml:"jushita_jms_topics_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误信息
    ResultMessage   string `json:"result_message,omitempty" xml:"result_message,omitempty"`
    // topic列表
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    // 错误码
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}
