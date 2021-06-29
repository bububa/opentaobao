package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认到店 API返回值 
alibaba.alihealth.booking.reserve.checkin

消费医疗统一预约平台，ISV 确认到店
*/
type AlibabaAlihealthBookingReserveCheckinAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBookingReserveCheckinResponse
}

// 确认到店 成功返回结果
type AlibabaAlihealthBookingReserveCheckinResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_checkin_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
