package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveConfirmAPIResponse
确认预约 API返回值
alibaba.alihealth.booking.reserve.confirm

确认预约 */
type AlibabaAlihealthBookingReserveConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBookingReserveConfirmAPIResponseModel
}

// AlibabaAlihealthBookingReserveConfirmAPIResponseModel is 确认预约 成功返回结果
type AlibabaAlihealthBookingReserveConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
