package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse 预约单状态查询 API返回值
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
type CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupAppointmentOrderStatusAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupAppointmentOrderStatusAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupAppointmentOrderStatusAPIResponseModel is 预约单状态查询 成功返回结果
type CainiaoGlobalImPickupAppointmentOrderStatusAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_appointment_order_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupAppointmentOrderStatusAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupAppointmentOrderStatusAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse)
	},
}

// GetCainiaoGlobalImPickupAppointmentOrderStatusAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse
func GetCainiaoGlobalImPickupAppointmentOrderStatusAPIResponse() *CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse {
	return poolCainiaoGlobalImPickupAppointmentOrderStatusAPIResponse.Get().(*CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse)
}

// ReleaseCainiaoGlobalImPickupAppointmentOrderStatusAPIResponse 将 CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupAppointmentOrderStatusAPIResponse(v *CainiaoGlobalImPickupAppointmentOrderStatusAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupAppointmentOrderStatusAPIResponse.Put(v)
}
