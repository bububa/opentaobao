package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoglobalimpickupappointmentorderinfocreateAPIResponse 首公里揽收-预约单创建 API返回值
// cainiao.global.im.pickup.appointment.order.info.create
//
// 预约单创建
type CainiaoglobalimpickupappointmentorderinfocreateAPIResponse struct {
	model.CommonResponse
	CainiaoglobalimpickupappointmentorderinfocreateAPIResponseModel
}

// CainiaoglobalimpickupappointmentorderinfocreateAPIResponseModel is 首公里揽收-预约单创建 成功返回结果
type CainiaoglobalimpickupappointmentorderinfocreateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_appointment_order_info_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
