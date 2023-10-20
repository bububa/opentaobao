package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupappointmentordercancelAPIResponse 首公里揽收-取消预约单 API返回值
// cainiao.global.im.pickup.appointment.order.cancel
//
// 首公里揽收-取消预约单创建
type CainiaoglobalimpickupappointmentordercancelAPIResponse struct {
	model.CommonResponse
	CainiaoglobalimpickupappointmentordercancelAPIResponseModel
}

// CainiaoglobalimpickupappointmentordercancelAPIResponseModel is 首公里揽收-取消预约单 成功返回结果
type CainiaoglobalimpickupappointmentordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_appointment_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
