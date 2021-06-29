package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通知改期结果 APIResponse
alibaba.alihealth.examination.reserve.modify.notify

体检状态为改期中，服务上通知健康是否改期成功
*/
type AlibabaAlihealthExaminationReserveModifyNotifyAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReserveModifyNotifyResponse
}

type AlibabaAlihealthExaminationReserveModifyNotifyResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_modify_notify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
