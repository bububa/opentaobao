package mos

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取指定流程上下文参数 API返回值 
alibaba.mosflow.work.queryvariables

业务查询指定流程上下文内容
*/
type AlibabaMosflowWorkQueryvariablesAPIResponse struct {
    model.CommonResponse
    AlibabaMosflowWorkQueryvariablesAPIResponseModel
}

// 获取指定流程上下文参数 成功返回结果
type AlibabaMosflowWorkQueryvariablesAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_mosflow_work_queryvariables_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *MultiResult `json:"result,omitempty" xml:"result,omitempty"`
}
