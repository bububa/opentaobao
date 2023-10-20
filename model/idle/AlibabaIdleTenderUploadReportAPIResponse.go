package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidletenderuploadreportAPIResponse 服务商上传验货报告同步给闲鱼 API返回值
// alibaba.idle.tender.upload.report
//
// 服务商上传验货报告同步给闲鱼
type AlibabaidletenderuploadreportAPIResponse struct {
	model.CommonResponse
	AlibabaidletenderuploadreportAPIResponseModel
}

// AlibabaidletenderuploadreportAPIResponseModel is 服务商上传验货报告同步给闲鱼 成功返回结果
type AlibabaidletenderuploadreportAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_tender_upload_report_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *IdleCommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
