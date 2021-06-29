package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV调TOP主动发起改期信息 APIResponse
alibaba.alihealth.examination.reserve.isv.modify

体检机构对接_ISV发起体检套餐改期
*/
type AlibabaAlihealthExaminationReserveIsvModifyAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReserveIsvModifyResponse
}

type AlibabaAlihealthExaminationReserveIsvModifyResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_isv_modify_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
