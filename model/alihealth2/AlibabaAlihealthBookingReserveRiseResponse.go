package alihealth2

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV 新增/修改复诊预约信息 APIResponse
alibaba.alihealth.booking.reserve.rise

ISV 新增/修改复诊预约信息
*/
type AlibabaAlihealthBookingReserveRiseAPIResponse struct {
    model.CommonResponse
    AlibabaAlihealthBookingReserveRiseResponse
}

type AlibabaAlihealthBookingReserveRiseResponse struct {
    XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_rise_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // alinkappserver系统返回的通用结果类
    
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
