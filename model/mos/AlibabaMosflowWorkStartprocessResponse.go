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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mosflow_work_startprocess_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 响应参数
    
    Data   string `json:"data,omitempty" xml:"