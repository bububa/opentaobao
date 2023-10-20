package wlbimports

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponseModel is 大包物流轨迹查询 成功返回结果
type CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_logistics_trajectory_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse)
	},
}

// GetCainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse
func GetCainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse() *CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse {
	return poolCainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse.Get().(*CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse)
}

// ReleaseCainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse 将 CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse(v *CainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagLogisticsTrajectoryAPIResponse.Put(v)
}
