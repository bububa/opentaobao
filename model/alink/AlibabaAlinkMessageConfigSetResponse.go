package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
消息提醒开关 APIResponse
alibaba.alink.message.config.set

阿里智能消息开关
*/
type AlibabaAlinkMessageConfigSetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkMessageConfigSetResponse `json:"alibaba_alink_message_config_set_response,omitempty"` 
    AlibabaAlinkMessageConfigSetResponse
}

/* model for simplify = false
type AlibabaAlinkMessageConfigSetResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkMessageConfigSetResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
