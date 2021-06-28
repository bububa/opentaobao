package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询消息列表 APIResponse
alibaba.alink.message.history.list

查询消息列表
*/
type AlibabaAlinkMessageHistoryListAPIResponse struct {
    model.CommonResponse
    AlibabaAlinkMessageHistoryListResponse
}

type AlibabaAlinkMessageHistoryListResponse struct {
    XMLName xml.Name `xml:"alibaba_alink_message_history_list_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
