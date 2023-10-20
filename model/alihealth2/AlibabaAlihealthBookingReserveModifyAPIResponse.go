package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthBookingReserveModifyAPIResponse 修改预约 API返回值
// alibaba.alihealth.booking.reserve.modify
//
// 消费医疗统一预约平台，取消预约
type AlibabaAlihealthBookingReserveModifyAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthBookingReserveModifyAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthBookingReserveModifyAPIResponseModel).Reset()
}

// AlibabaAlihealthBookingReserveModifyAPIResponseModel is 修改预约 成功返回结果
type AlibabaAlihealthBookingReserveModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_booking_reserve_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// alinkappserver系统返回的通用结果类
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthBookingReserveModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthBookingReserveModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthBookingReserveModifyAPIResponse)
	},
}

// GetAlibabaAlihealthBookingReserveModifyAPIResponse 从 sync.Pool 获取 AlibabaAlihealthBookingReserveModifyAPIResponse
func GetAlibabaAlihealthBookingReserveModifyAPIResponse() *AlibabaAlihealthBookingReserveModifyAPIResponse {
	return poolAlibabaAlihealthBookingReserveModifyAPIResponse.Get().(*AlibabaAlihealthBookingReserveModifyAPIResponse)
}

// ReleaseAlibabaAlihealthBookingReserveModifyAPIResponse 将 AlibabaAlihealthBookingReserveModifyAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthBookingReserveModifyAPIResponse(v *AlibabaAlihealthBookingReserveModifyAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthBookingReserveModifyAPIResponse.Put(v)
}
