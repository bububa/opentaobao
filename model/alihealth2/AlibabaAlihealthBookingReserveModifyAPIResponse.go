package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBookingReserveModifyAPIResponse
修改预约 API返回值
alibaba.alihealth.booking.reserve.modify

消费医疗统一预约平台，取消预约 */
type AlibabaAlihealthBookingReserveModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBookingReserveModifyAPIResponseModel
}

// AlibabaAlihealthBookingReserveModifyAPIResponseModel is 修改预约 成功返回结果
type AlibabaAlihealthBookingReserveModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
