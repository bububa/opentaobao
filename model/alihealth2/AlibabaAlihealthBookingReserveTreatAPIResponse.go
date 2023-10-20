package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveTreatAPIResponse 确认就诊 API返回值
// alibaba.alihealth.booking.reserve.treat
//
// 统一预约平台，ISV确认就诊
type AlibabaAlihealthBookingReserveTreatAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBookingReserveTreatAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveTreatAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthBookingReserveTreatAPIResponseModel).Reset()
}

// AlibabaAlihealthBookingReserveTreatAPIResponseModel is 确认就诊 成功返回结果
type AlibabaAlihealthBookingReserveTreatAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_treat_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveTreatAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthBookingReserveTreatAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthBookingReserveTreatAPIResponse)
	},
}

// GetAlibabaAlihealthBookingReserveTreatAPIResponse 从 sync.Pool 获取 AlibabaAlihealthBookingReserveTreatAPIResponse
func GetAlibabaAlihealthBookingReserveTreatAPIResponse() *AlibabaAlihealthBookingReserveTreatAPIResponse {
	return poolAlibabaAlihealthBookingReserveTreatAPIResponse.Get().(*AlibabaAlihealthBookingReserveTreatAPIResponse)
}

// ReleaseAlibabaAlihealthBookingReserveTreatAPIResponse 将 AlibabaAlihealthBookingReserveTreatAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveTreatAPIResponse(v *AlibabaAlihealthBookingReserveTreatAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveTreatAPIResponse.Put(v)
}
