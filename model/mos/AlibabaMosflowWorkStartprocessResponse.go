package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发起流程 APIResponse
alibaba.mosflow.work.startprocess

业务发起流程审批
*/
type AlibabaMosflowWorkStartprocessAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMosflowWorkStartprocessResponse `json:"alibaba_mosflow_work_startprocess_response,omitempty"` 
    AlibabaMosflowWorkStartprocessResponse
}

/* model for simplify = false
type AlibabaMosflowWorkStartprocessResponse struct {

    // 响应参数
    
    Data   string `json:"data,omitempty"`
    

    // 操作是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 异常信息
    
    ResultMessage   string `json:"result_message,omitempty"`
    

    // 异常Code
    
    ResultCode   string `json:"result_code,omitempty"`
    

}
*/

type AlibabaMosflowWorkStartprocessResponse struct {

    // 响应参数
    Data   string `json:"data,omitempty"`

    // 操作是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 异常信息
    ResultMessage   string `json:"result_message,omitempty"`

    // 异常Code
    ResultCode   string `json:"result_code,omitempty"`

}
