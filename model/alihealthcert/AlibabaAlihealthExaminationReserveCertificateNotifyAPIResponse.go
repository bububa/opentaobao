package alihealthcert

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse 健康证服务商预约结果通知阿里健康 API返回值
// alibaba.alihealth.examination.reserve.certificate.notify
//
// 当ISV执行完健康证预约成功之后， 调用通知阿里健康
type AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponseModel).Reset()
}

// AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponseModel is 健康证服务商预约结果通知阿里健康 成功返回结果
type AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_certificate_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse)
	},
}

// GetAlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse
func GetAlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse() *AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse {
	return poolAlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse.Get().(*AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse)
}

// ReleaseAlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse 将 AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse(v *AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse.Put(v)
}
