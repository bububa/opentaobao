package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
确认到店 APIResponse
alibaba.alihealth.booking.reserve.checkin

消费医疗统一预约平台，ISV 确认到店
*/
type AlibabaAlihealthBookingReserveCheckinAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBookingReserveCheckinResponse
}

type AlibabaAlihealthBookingReserveCheckinResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_checkin_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
