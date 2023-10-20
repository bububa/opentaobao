package tmallgeniescp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanfeedbackoemuploadAPIResponse 14-供应商反馈（OEM）同步接口 API返回值
// alibaba.tmallgenie.scp.plan.feedback.oem.upload
//
// 供应商反馈（OEM）同步接口
type AlibabatmallgeniescpplanfeedbackoemuploadAPIResponse struct {
	model.CommonResponse
	AlibabatmallgeniescpplanfeedbackoemuploadAPIResponseModel
}

// AlibabatmallgeniescpplanfeedbackoemuploadAPIResponseModel is 14-供应商反馈（OEM）同步接口 成功返回结果
type AlibabatmallgeniescpplanfeedbackoemuploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tmallgenie_scp_plan_feedback_oem_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回对象封装
	Result *DataResult `json:"result,omitempty" xml:"result,omitempty"`
}
