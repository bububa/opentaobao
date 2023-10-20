package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse 预约单差异明细查询 API返回值
// cainiao.global.im.pickup.appointment.order.difference.detail
//
// 预约单差异明细查询
type CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponseModel is 预约单差异明细查询 成功返回结果
type CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_appointment_order_difference_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse)
	},
}

// GetCainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse
func GetCainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse() *CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse {
	return poolCainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse.Get().(*CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse)
}

// ReleaseCainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse 将 CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse(v *CainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupAppointmentOrderDifferenceDetailAPIResponse.Put(v)
}
