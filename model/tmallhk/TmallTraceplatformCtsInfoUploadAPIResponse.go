package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformCtsInfoUploadAPIResponse CTS提交溯源信息 API返回值
// tmall.traceplatform.cts.info.upload
//
// cts上传溯源信息
type TmallTraceplatformCtsInfoUploadAPIResponse struct {
	model.CommonResponse
	TmallTraceplatformCtsInfoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraceplatformCtsInfoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraceplatformCtsInfoUploadAPIResponseModel).Reset()
}

// TmallTraceplatformCtsInfoUploadAPIResponseModel is CTS提交溯源信息 成功返回结果
type TmallTraceplatformCtsInfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_cts_info_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraceplatformCtsInfoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTraceplatformCtsInfoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraceplatformCtsInfoUploadAPIResponse)
	},
}

// GetTmallTraceplatformCtsInfoUploadAPIResponse 从 sync.Pool 获取 TmallTraceplatformCtsInfoUploadAPIResponse
func GetTmallTraceplatformCtsInfoUploadAPIResponse() *TmallTraceplatformCtsInfoUploadAPIResponse {
	return poolTmallTraceplatformCtsInfoUploadAPIResponse.Get().(*TmallTraceplatformCtsInfoUploadAPIResponse)
}

// ReleaseTmallTraceplatformCtsInfoUploadAPIResponse 将 TmallTraceplatformCtsInfoUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallTraceplatformCtsInfoUploadAPIResponse(v *TmallTraceplatformCtsInfoUploadAPIResponse) {
	v.Reset()
	poolTmallTraceplatformCtsInfoUploadAPIResponse.Put(v)
}
