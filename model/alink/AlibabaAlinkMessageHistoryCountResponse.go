package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询消息总数 APIResponse
alibaba.alink.message.history.count

查询消息总数
*/
type AlibabaAlinkMessageHistoryCountAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkMessageHistoryCountResponse `json:"alibaba_alink_message_history_count_response,omitempty"` 
    AlibabaAlinkMessageHistoryCountResponse
}

/* model for simplify = false
type AlibabaAlinkMessageHistoryCountResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkMessageHistoryCountResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
