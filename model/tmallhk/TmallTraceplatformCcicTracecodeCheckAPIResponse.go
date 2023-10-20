package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformCcicTracecodeCheckAPIResponse ccic校验溯源码 API返回值
// tmall.traceplatform.ccic.tracecode.check
//
// 天猫国际溯源业务，需要将溯源码校验的功能输出到ccic官方主页中以增强溯源码的可信度，故需要提供api给ccic使用以校验溯源码的正确性。
type TmallTraceplatformCcicTracecodeCheckAPIResponse struct {
	model.CommonResponse
	TmallTraceplatformCcicTracecodeCheckAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraceplatformCcicTracecodeCheckAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraceplatformCcicTracecodeCheckAPIResponseModel).Reset()
}

// TmallTraceplatformCcicTracecodeCheckAPIResponseModel is ccic校验溯源码 成功返回结果
type TmallTraceplatformCcicTracecodeCheckAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_ccic_tracecode_check_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraceplatformCcicTracecodeCheckAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTraceplatformCcicTracecodeCheckAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraceplatformCcicTracecodeCheckAPIResponse)
	},
}

// GetTmallTraceplatformCcicTracecodeCheckAPIResponse 从 sync.Pool 获取 TmallTraceplatformCcicTracecodeCheckAPIResponse
func GetTmallTraceplatformCcicTracecodeCheckAPIResponse() *TmallTraceplatformCcicTracecodeCheckAPIResponse {
	return poolTmallTraceplatformCcicTracecodeCheckAPIResponse.Get().(*TmallTraceplatformCcicTracecodeCheckAPIResponse)
}

// ReleaseTmallTraceplatformCcicTracecodeCheckAPIResponse 将 TmallTraceplatformCcicTracecodeCheckAPIResponse 保存到 sync.Pool
func ReleaseTmallTraceplatformCcicTracecodeCheckAPIResponse(v *TmallTraceplatformCcicTracecodeCheckAPIResponse) {
	v.Reset()
	poolTmallTraceplatformCcicTracecodeCheckAPIResponse.Put(v)
}
