package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlereportresultuploadAPIResponse 服务商上传验货报告 API返回值
// alibaba.idle.report.result.upload
//
// 服务商上传验货报告
type AlibabaidlereportresultuploadAPIResponse struct {
	model.CommonResponse
	AlibabaidlereportresultuploadAPIResponseModel
}

// AlibabaidlereportresultuploadAPIResponseModel is 服务商上传验货报告 成功返回结果
type AlibabaidlereportresultuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_report_result_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
