package jms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据用户nick获取开通的消息列表 APIResponse
taobao.jushita.jms.topics.get

根据用户nick获取开通的消息列表
*/
type TaobaoJushitaJmsTopicsGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"jushita_jms_topics_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误信息
    
    ResultMessage   string `json:"result_message,omitempty" xml:"