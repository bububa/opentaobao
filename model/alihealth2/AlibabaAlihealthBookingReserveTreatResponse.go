package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认就诊 API返回值 
alibaba.alihealth.booking.reserve.treat

统一预约平台，ISV确认就诊
*/
type AlibabaAlihealthBookingReserveTreatAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBookingReserveTreatResponse
}

// 确认就诊 成功返回结果
type AlibabaAlihealthBookingReserveTreatResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_treat_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
