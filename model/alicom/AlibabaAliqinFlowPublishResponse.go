package alicom

import (
    "github.com/bububa/opentaobao/model"
)

/* 
流量发放(用户id) APIResponse
alibaba.aliqin.flow.publish

阿里通信流量下发功能，允许用户补发
*/
type AlibabaAliqinFlowPublishAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaAliqinFlowPublishResponse `json:"alibaba_aliqin_flow_publish_response,omitempty"` 
    AlibabaAliqinFlowPublishResponse
}

/* model for simplify = false
type AlibabaAliqinFlowPublishResponse struct {

    // true为成功，其他为失败
    
    Value   string `json:"value,omitempty"`
    

}
*/

type AlibabaAliqinFlowPublishResponse struct {

    // true为成功，其他为失败
    Value   string `json:"value,omitempty"`

}
