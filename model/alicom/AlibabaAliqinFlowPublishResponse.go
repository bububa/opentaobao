package alicom

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
流量发放(用户id) APIResponse
alibaba.aliqin.flow.publish

阿里通信流量下发功能，允许用户补发
*/
type AlibabaAliqinFlowPublishAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_aliqin_flow_publish_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // true为成功，其他为失败
    
    Value   string `json:"value,omitempty" xml:"