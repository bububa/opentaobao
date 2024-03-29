package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse 供应商校准后的配额同步 API返回值
// alibaba.tmallgenie.scp.plan.correct.supplier.quote.upload
//
// 供应商校准后的配额同步
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponseModel is 供应商校准后的配额同步 成功返回结果
type AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_correct_supplier_quote_upload_response"`
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
func (m *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DataList = m.DataList[:0]
	m.ResultMsg = ""
	m.TraceId = ""
	m.ResultCode = ""
}

var poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse
func GetAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse() *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse 将 AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse(v *AlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanCorrectSupplierQuoteUploadAPIResponse.Put(v)
}
