package examination

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationinvoiceinfonotifyAPIResponse 体检机构同步发票信息给阿里健康 API返回值
// alibaba.alihealth.examination.invoice.info.notify
//
// 体检机构向阿里健康同步发票信息
type AlibabaalihealthexaminationinvoiceinfonotifyAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationinvoiceinfonotifyAPIResponseModel
}

// AlibabaalihealthexaminationinvoiceinfonotifyAPIResponseModel is 体检机构同步发票信息给阿里健康 成功返回结果
type AlibabaalihealthexaminationinvoiceinfonotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_invoice_info_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
