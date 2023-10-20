package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanforecastoemuploadAPIResponse 16-供应商预测（OEM-成品）回传接口 API返回值
// alibaba.tmallgenie.scp.plan.forecast.oem.upload
//
// 供应商预测（OEM-成品）回传接口
type AlibabatmallgeniescpplanforecastoemuploadAPIResponse struct {
	model.CommonResponse
	AlibabatmallgeniescpplanforecastoemuploadAPIResponseModel
}

// AlibabatmallgeniescpplanforecastoemuploadAPIResponseModel is 16-供应商预测（OEM-成品）回传接口 成功返回结果
type AlibabatmallgeniescpplanforecastoemuploadAPIResponseModel struct {
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
