package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通知改期结果 API返回值 
alibaba.alihealth.examination.reserve.modify.notify

体检状态为改期中，服务上通知健康是否改期成功
*/
type AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReserveModifyNotifyResponse
}

// 通知改期结果 成功返回结果
type AlibabaAlihealthExaminationReserveModifyNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_modify_notify_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
