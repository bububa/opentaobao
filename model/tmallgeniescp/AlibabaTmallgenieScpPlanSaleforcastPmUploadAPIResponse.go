package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse 18-销售预测数量（产管）回传接口 API返回值
// alibaba.tmallgenie.scp.plan.saleforcast.pm.upload
//
// 销售预测数量（产管）回传接口
type AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponseModel is 18-销售预测数量（产管）回传接口 成功返回结果
type AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_saleforcast_pm_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回描述
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 返回码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.TraceId = ""
	m.ResultCode = ""
}

var poolAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse
func GetAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse() *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse 将 AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse(v *AlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanSaleforcastPmUploadAPIResponse.Put(v)
}
