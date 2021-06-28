package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定流程上下文参数 APIResponse
alibaba.mosflow.work.queryvariables

业务查询指定流程上下文内容
*/
type AlibabaMosflowWorkQueryvariablesAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_mosflow_work_queryvariables_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *MultiResult `json:"result,omitempty" xml:"