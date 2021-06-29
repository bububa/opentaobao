package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询消息总数 APIResponse
alibaba.alink.message.history.count

查询消息总数
*/
type AlibabaAlinkMessageHistoryCountAPIResponse struct {
    model.CommonResponse
    AlibabaAlinkMessageHistoryCountResponse
}

type AlibabaAlinkMessageHistoryCountResponse struct {
    XMLName xml.Name `xml:"alibaba_alink_message_history_count_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
