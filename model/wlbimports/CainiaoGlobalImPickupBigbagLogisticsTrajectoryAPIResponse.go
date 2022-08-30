package wlbimports

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse 大包物流轨迹查询 API返回值
// cainiao.global.im.pickup.bigbag.logistics.trajectory
//
// 大包物流轨迹查询
type CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponseModel
}

// CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponseModel is 大包物流轨迹查询 成功返回结果
type CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_logistics_trajectory_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
