package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveCheckinAPIResponse 确认到店 API返回值
// alibaba.alihealth.booking.reserve.checkin
//
// 消费医疗统一预约平台，ISV 确认到店
type AlibabaAlihealthBookingReserveCheckinAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBookingReserveCheckinAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveCheckinAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthBookingReserveCheckinAPIResponseModel).Reset()
}

// AlibabaAlihealthBookingReserveCheckinAPIResponseModel is 确认到店 成功返回结果
type AlibabaAlihealthBookingReserveCheckinAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_checkin_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveCheckinAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthBookingReserveCheckinAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthBookingReserveCheckinAPIResponse)
	},
}

// GetAlibabaAlihealthBookingReserveCheckinAPIResponse 从 sync.Pool 获取 AlibabaAlihealthBookingReserveCheckinAPIResponse
func GetAlibabaAlihealthBookingReserveCheckinAPIResponse() *AlibabaAlihealthBookingReserveCheckinAPIResponse {
	return poolAlibabaAlihealthBookingReserveCheckinAPIResponse.Get().(*AlibabaAlihealthBookingReserveCheckinAPIResponse)
}

// ReleaseAlibabaAlihealthBookingReserveCheckinAPIResponse 将 AlibabaAlihealthBookingReserveCheckinAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveCheckinAPIResponse(v *AlibabaAlihealthBookingReserveCheckinAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveCheckinAPIResponse.Put(v)
}
