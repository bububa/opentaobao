package examination

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
体检机构对接_预约取消 API返回值 
alibaba.alihealth.examination.reserve.cancel

体检机构对接_体检取消
*/
type AlibabaAlihealthExaminationReserveCancelAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthExaminationReserveCancelAPIResponseModel
}

// 体检机构对接_预约取消 成功返回结果
type AlibabaAlihealthExaminationReserveCancelAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_alihealth_examination_reserve_cancel_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 200:取消成功，状态变为已取消;800:取消失败，状态变为取消预约中；700：取消失败状态从取消预约中变为已到检；600取消失败，状态从取消预约中变为已预约
    ResponseCode   string `json:"response_code,omitempty" xml:"response_code,omitempty"`
    // 返回状态码不是600或者700时，不需要返回此字段
    RevisionInfo   *RevisionInfo `json:"revision_info,omitempty" xml:"revision_info,omitempty"`
    // 返回结果描述
    Message   string `json:"message,omitempty" xml:"message,omitempty"`
}
