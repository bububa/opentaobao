package alink

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询消息开关配置列表 APIResponse
alibaba.alink.message.config.list

阿里智能获取消息开关配置列表
*/
type AlibabaAlinkMessageConfigListAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_alink_message_config_list_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *TopServiceResult `json:"result,omitempty" xml:"