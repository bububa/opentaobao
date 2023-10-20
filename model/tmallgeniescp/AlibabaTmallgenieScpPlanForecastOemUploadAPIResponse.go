package tmallgeniescp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse 16-供应商预测（OEM-成品）回传接口 API返回值
// alibaba.tmallgenie.scp.plan.forecast.oem.upload
//
// 供应商预测（OEM-成品）回传接口
type AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse struct {
	model.CommonResponse
	AlibabaTmallgenieScpPlanForecastOemUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaTmallgenieScpPlanForecastOemUploadAPIResponseModel).Reset()
}

// AlibabaTmallgenieScpPlanForecastOemUploadAPIResponseModel is 16-供应商预测（OEM-成品）回传接口 成功返回结果
type AlibabaTmallgenieScpPlanForecastOemUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_forecast_oem_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 参数msg
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求唯一ID
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 参数code
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaTmallgenieScpPlanForecastOemUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.TraceId = ""
	m.ResultCode = ""
}

var poolAlibabaTmallgenieScpPlanForecastOemUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse)
	},
}

// GetAlibabaTmallgenieScpPlanForecastOemUploadAPIResponse 从 sync.Pool 获取 AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse
func GetAlibabaTmallgenieScpPlanForecastOemUploadAPIResponse() *AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse {
	return poolAlibabaTmallgenieScpPlanForecastOemUploadAPIResponse.Get().(*AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse)
}

// ReleaseAlibabaTmallgenieScpPlanForecastOemUploadAPIResponse 将 AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaTmallgenieScpPlanForecastOemUploadAPIResponse(v *AlibabaTmallgenieScpPlanForecastOemUploadAPIResponse) {
	v.Reset()
	poolAlibabaTmallgenieScpPlanForecastOemUploadAPIResponse.Put(v)
}
