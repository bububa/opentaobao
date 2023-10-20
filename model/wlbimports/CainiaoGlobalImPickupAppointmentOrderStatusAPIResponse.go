package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupappointmentorderstatusAPIResponse 预约单状态查询 API返回值
// cainiao.global.im.pickup.appointment.order.status
//
// 预约单状态查询
type CainiaoglobalimpickupappointmentorderstatusAPIResponse struct {
	model.CommonResponse
	CainiaoglobalimpickupappointmentorderstatusAPIResponseModel
}

// CainiaoglobalimpickupappointmentorderstatusAPIResponseModel is 预约单状态查询 成功返回结果
type CainiaoglobalimpickupappointmentorderstatusAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_appointment_order_status_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
