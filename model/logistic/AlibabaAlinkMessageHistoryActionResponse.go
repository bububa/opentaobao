package logistic

import (
    "github.com/bububa/opentaobao/model"
)

/* 
操作历史消息 APIResponse
alibaba.alink.message.history.action

阿里智能操作历史消息
*/
type AlibabaAlinkMessageHistoryActionAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkMessageHistoryActionResponse `json:"alibaba_alink_message_history_action_response,omitempty"` 
    AlibabaAlinkMessageHistoryActionResponse
}

/* model for simplify = false
type AlibabaAlinkMessageHistoryActionResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkMessageHistoryActionResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
