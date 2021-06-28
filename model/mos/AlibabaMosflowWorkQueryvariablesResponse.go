package mos

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取指定流程上下文参数 APIResponse
alibaba.mosflow.work.queryvariables

业务查询指定流程上下文内容
*/
type AlibabaMosflowWorkQueryvariablesAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMosflowWorkQueryvariablesResponse `json:"alibaba_mosflow_work_queryvariables_response,omitempty"` 
    AlibabaMosflowWorkQueryvariablesResponse
}

/* model for simplify = false
type AlibabaMosflowWorkQueryvariablesResponse struct {

    // result
    
    Result  *struct {
        MultiResult  *MultiResult `json:"multi_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaMosflowWorkQueryvariablesResponse struct {

    // result
    Result   *MultiResult `json:"result,omitempty"`

}
