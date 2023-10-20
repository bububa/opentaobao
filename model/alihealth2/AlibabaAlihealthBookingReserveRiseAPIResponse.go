package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveRiseAPIResponse ISV 新增/修改复诊预约信息 API返回值
// alibaba.alihealth.booking.reserve.rise
//
// ISV 新增/修改复诊预约信息
type AlibabaAlihealthBookingReserveRiseAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBookingReserveRiseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveRiseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthBookingReserveRiseAPIResponseModel).Reset()
}

// AlibabaAlihealthBookingReserveRiseAPIResponseModel is ISV 新增/修改复诊预约信息 成功返回结果
type AlibabaAlihealthBookingReserveRiseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_rise_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveRiseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthBookingReserveRiseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthBookingReserveRiseAPIResponse)
	},
}

// GetAlibabaAlihealthBookingReserveRiseAPIResponse 从 sync.Pool 获取 AlibabaAlihealthBookingReserveRiseAPIResponse
func GetAlibabaAlihealthBookingReserveRiseAPIResponse() *AlibabaAlihealthBookingReserveRiseAPIResponse {
	return poolAlibabaAlihealthBookingReserveRiseAPIResponse.Get().(*AlibabaAlihealthBookingReserveRiseAPIResponse)
}

// ReleaseAlibabaAlihealthBookingReserveRiseAPIResponse 将 AlibabaAlihealthBookingReserveRiseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveRiseAPIResponse(v *AlibabaAlihealthBookingReserveRiseAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveRiseAPIResponse.Put(v)
}
