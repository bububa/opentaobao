package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
发起流程 APIResponse
alibaba.mosflow.work.startprocess

业务发起流程审批
*/
type AlibabaMosflowWorkStartprocessAPIResponse struct {
    model.CommonResponse
    AlibabaMosflowWorkStartprocessResponse
}

type AlibabaMosflowWorkStartprocessResponse struct {
    XMLName xml.Name `xml:"alibaba_mosflow_work_startprocess_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 响应参数
    
    Data   string `json:"data,omitempty" xml:"data,omitempty"`

    
    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 异常信息
    
    ResultMessage   string `json:"result_message,omitempty" xml:"result_message,omitempty"`

    
    // 异常Code
    
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`

    
}
