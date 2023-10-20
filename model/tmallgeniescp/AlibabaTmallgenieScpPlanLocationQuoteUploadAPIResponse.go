package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse 9.2-同步地点配额 API返回值
// alibaba.tmallgenie.scp.plan.location.quote.upload
//
// 同步地点配额
type AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponseModel is 9.2-同步地点配额 成功返回结果
type AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_location_quote_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 对象列表
	DataList []ChannelQuotaDto `json:"data_list,omitempty" xml:"data_list>channel_quota_dto,omitempty"`
	// 返回描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataList = m.DataList[:0]
	m.ResultMsg = ""
	m.TraceId = ""
	m.ResultCode = ""
}

var poolAlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse
func GetAlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse() *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse 将 AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse(v *AlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanLocationQuoteUploadAPIResponse.Put(v)
}
