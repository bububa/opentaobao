package examination

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse 体检机构同步发票信息给阿里健康 API返回值
// alibaba.alihealth.examination.invoice.info.notify
//
// 体检机构向阿里健康同步发票信息
type AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponseModel is 体检机构同步发票信息给阿里健康 成功返回结果
type AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_invoice_info_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse
func GetAlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse() *AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse {
	return poolAlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse.Get().(*AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse 将 AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse(v *AlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationInvoiceInfoNotifyAPIResponse.Put(v)
}
