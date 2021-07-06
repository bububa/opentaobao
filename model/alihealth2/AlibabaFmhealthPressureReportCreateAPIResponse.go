package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFmhealthPressureReportCreateAPIResponse 血压报告接口 API返回值
// alibaba.fmhealth.pressure.report.create
//
// 生成用户血压测量报告
type AlibabaFmhealthPressureReportCreateAPIResponse struct {
	model.CommonResponse
	AlibabaFmhealthPressureReportCreateAPIResponseModel
}

// AlibabaFmhealthPressureReportCreateAPIResponseModel is 血压报告接口 成功返回结果
type AlibabaFmhealthPressureReportCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fmhealth_pressure_report_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// msgCode
	MsgCode int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// success
	Status bool `json:"status,omitempty" xml:"status,omitempty"`
}
