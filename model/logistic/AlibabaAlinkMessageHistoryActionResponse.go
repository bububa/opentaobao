package logistic

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
操作历史消息 APIResponse
alibaba.alink.message.history.action

阿里智能操作历史消息
*/
type AlibabaAlinkMessageHistoryActionAPIResponse struct {
    model.CommonResponse
    AlibabaAlinkMessageHistoryActionResponse
}

type AlibabaAlinkMessageHistoryActionResponse struct {
    XMLName xml.Name `xml:"alibaba_alink_message_history_action_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
