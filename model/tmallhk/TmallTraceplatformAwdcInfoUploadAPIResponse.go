package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformAwdcInfoUploadAPIResponse AWDC提交溯源信息 API返回值
// tmall.traceplatform.awdc.info.upload
//
// 天猫溯源-AWDC-上传溯源信息
type TmallTraceplatformAwdcInfoUploadAPIResponse struct {
	model.CommonResponse
	TmallTraceplatformAwdcInfoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraceplatformAwdcInfoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraceplatformAwdcInfoUploadAPIResponseModel).Reset()
}

// TmallTraceplatformAwdcInfoUploadAPIResponseModel is AWDC提交溯源信息 成功返回结果
type TmallTraceplatformAwdcInfoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traceplatform_awdc_info_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraceplatformAwdcInfoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallTraceplatformAwdcInfoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraceplatformAwdcInfoUploadAPIResponse)
	},
}

// GetTmallTraceplatformAwdcInfoUploadAPIResponse 从 sync.Pool 获取 TmallTraceplatformAwdcInfoUploadAPIResponse
func GetTmallTraceplatformAwdcInfoUploadAPIResponse() *TmallTraceplatformAwdcInfoUploadAPIResponse {
	return poolTmallTraceplatformAwdcInfoUploadAPIResponse.Get().(*TmallTraceplatformAwdcInfoUploadAPIResponse)
}

// ReleaseTmallTraceplatformAwdcInfoUploadAPIResponse 将 TmallTraceplatformAwdcInfoUploadAPIResponse 保存到 sync.Pool
func ReleaseTmallTraceplatformAwdcInfoUploadAPIResponse(v *TmallTraceplatformAwdcInfoUploadAPIResponse) {
	v.Reset()
	poolTmallTraceplatformAwdcInfoUploadAPIResponse.Put(v)
}
