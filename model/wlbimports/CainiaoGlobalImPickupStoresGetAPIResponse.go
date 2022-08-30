package wlbimports

import (
	"encoding/xml"

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

// CainiaoGlobalImPickupStoresGetAPIResponseModel is 首公里揽收-集货仓列表查询 成功返回结果
type CainiaoGlobalImPickupStoresGetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_stores_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应体
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}
