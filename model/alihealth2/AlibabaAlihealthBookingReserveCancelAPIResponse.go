package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveCancelAPIResponse 取消预约 API返回值
// alibaba.alihealth.booking.reserve.cancel
//
// 消费医疗统一预约平台，ISV取消预约
type AlibabaAlihealthBookingReserveCancelAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBookingReserveCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthBookingReserveCancelAPIResponseModel).Reset()
}

// AlibabaAlihealthBookingReserveCancelAPIResponseModel is 取消预约 成功返回结果
type AlibabaAlihealthBookingReserveCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthBookingReserveCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthBookingReserveCancelAPIResponse)
	},
}

// GetAlibabaAlihealthBookingReserveCancelAPIResponse 从 sync.Pool 获取 AlibabaAlihealthBookingReserveCancelAPIResponse
func GetAlibabaAlihealthBookingReserveCancelAPIResponse() *AlibabaAlihealthBookingReserveCancelAPIResponse {
	return poolAlibabaAlihealthBookingReserveCancelAPIResponse.Get().(*AlibabaAlihealthBookingReserveCancelAPIResponse)
}

// ReleaseAlibabaAlihealthBookingReserveCancelAPIResponse 将 AlibabaAlihealthBookingReserveCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveCancelAPIResponse(v *AlibabaAlihealthBookingReserveCancelAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveCancelAPIResponse.Put(v)
}
