package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse 9.1-同步渠道配额 API返回值
// alibaba.tmallgenie.scp.plan.channel.quote.upload
//
// 同步渠道配额
type AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponseModel is 9.1-同步渠道配额 成功返回结果
type AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_channel_quote_upload_response"`
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
func (m *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataList = m.DataList[:0]
	m.ResultMsg = ""
	m.TraceId = ""
	m.ResultCode = ""
}

var poolAlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse
func GetAlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse() *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse 将 AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse(v *AlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanChannelQuoteUploadAPIResponse.Put(v)
}
