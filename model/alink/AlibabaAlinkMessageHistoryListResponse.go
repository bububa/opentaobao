package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询消息列表 APIResponse
alibaba.alink.message.history.list

查询消息列表
*/
type AlibabaAlinkMessageHistoryListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkMessageHistoryListResponse `json:"alibaba_alink_message_history_list_response,omitempty"` 
    AlibabaAlinkMessageHistoryListResponse
}

/* model for simplify = false
type AlibabaAlinkMessageHistoryListResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkMessageHistoryListResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
