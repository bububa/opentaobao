package alihealthcert

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthexaminationreservecertificatenotifyAPIResponse 健康证服务商预约结果通知阿里健康 API返回值
// alibaba.alihealth.examination.reserve.certificate.notify
//
// 当ISV执行完健康证预约成功之后， 调用通知阿里健康
type AlibabaalihealthexaminationreservecertificatenotifyAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthexaminationreservecertificatenotifyAPIResponseModel
}

// AlibabaalihealthexaminationreservecertificatenotifyAPIResponseModel is 健康证服务商预约结果通知阿里健康 成功返回结果
type AlibabaalihealthexaminationreservecertificatenotifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_certificate_notify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
