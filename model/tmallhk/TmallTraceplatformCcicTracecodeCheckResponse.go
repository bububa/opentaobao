package tmallhk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
ccic校验溯源码 APIResponse
tmall.traceplatform.ccic.tracecode.check

天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。
*/
type TmallTraceplatformCcicTracecodeCheckAPIResponse struct {
    model.CommonResponse
    Response *TmallTraceplatformCcicTracecodeCheckResponse `json:"tmall_traceplatform_ccic_tracecode_check_response,omitempty"`
}

type TmallTraceplatformCcicTracecodeCheckResponse struct {

    // result
    Result   *DataResult `json:"result,omitempty"`

}