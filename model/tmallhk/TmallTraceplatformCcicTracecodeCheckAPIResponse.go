package tmallhk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformccictracecodecheckAPIResponse ccic校验溯源码 API返回值
// tmall.traceplatform.ccic.tracecode.check
//
// 天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。
type TmalltraceplatformccictracecodecheckAPIResponse struct {
	model.CommonResponse
	TmalltraceplatformccictracecodecheckAPIResponseModel
}

// TmalltraceplatformccictracecodecheckAPIResponseModel is ccic校验溯源码 成功返回结果
type TmalltraceplatformccictracecodecheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_ccic_tracecode_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
