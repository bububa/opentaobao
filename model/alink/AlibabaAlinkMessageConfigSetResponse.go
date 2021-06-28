package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
消息提醒开关 APIResponse
alibaba.alink.message.config.set

阿里智能消息开关
*/
type AlibabaAlinkMessageConfigSetAPIResponse struct {
    model.CommonResponse
    AlibabaAlinkMessageConfigSetResponse
}

type AlibabaAlinkMessageConfigSetResponse struct {
    XMLName xml.Name `xml:"alibaba_alink_message_config_set_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
