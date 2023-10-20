package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupStoresGetAPIResponse 首公里揽收-集货仓列表查询 API返回值
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
type CainiaoGlobalImPickupStoresGetAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupStoresGetAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupStoresGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupStoresGetAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupStoresGetAPIResponseModel is 首公里揽收-集货仓列表查询 成功返回结果
type CainiaoGlobalImPickupStoresGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_stores_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupStoresGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupStoresGetAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupStoresGetAPIResponse)
	},
}

// GetCainiaoGlobalImPickupStoresGetAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupStoresGetAPIResponse
func GetCainiaoGlobalImPickupStoresGetAPIResponse() *CainiaoGlobalImPickupStoresGetAPIResponse {
	return poolCainiaoGlobalImPickupStoresGetAPIResponse.Get().(*CainiaoGlobalImPickupStoresGetAPIResponse)
}

// ReleaseCainiaoGlobalImPickupStoresGetAPIResponse 将 CainiaoGlobalImPickupStoresGetAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupStoresGetAPIResponse(v *CainiaoGlobalImPickupStoresGetAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupStoresGetAPIResponse.Put(v)
}
