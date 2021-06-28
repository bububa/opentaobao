package alink

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询消息开关配置列表 APIResponse
alibaba.alink.message.config.list

阿里智能获取消息开关配置列表
*/
type AlibabaAlinkMessageConfigListAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAlinkMessageConfigListResponse `json:"alibaba_alink_message_config_list_response,omitempty"` 
    AlibabaAlinkMessageConfigListResponse
}

/* model for simplify = false
type AlibabaAlinkMessageConfigListResponse struct {

    // 结果
    
    Result  *struct {
        TopServiceResult  *TopServiceResult `json:"top_service_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaAlinkMessageConfigListResponse struct {

    // 结果
    Result   *TopServiceResult `json:"result,omitempty"`

}
