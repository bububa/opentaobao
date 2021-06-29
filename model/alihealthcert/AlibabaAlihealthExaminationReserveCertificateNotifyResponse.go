package alihealthcert

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
健康证服务商预约结果通知阿里健康 APIResponse
alibaba.alihealth.examination.reserve.certificate.notify

当ISV执行完健康证预约成功之后， 调用通知阿里健康
*/
type AlibabaAlihealthExaminationReserveCertificateNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReserveCertificateNotifyResponse
}

type AlibabaAlihealthExaminationReserveCertificateNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_certificate_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
