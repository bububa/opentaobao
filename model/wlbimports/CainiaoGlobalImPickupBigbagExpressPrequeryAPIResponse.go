package wlbimports

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse 首公里揽收-快递预查询服务 API返回值
// cainiao.global.im.pickup.bigbag.express.prequery
//
// 快递预查询服务
type CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse struct {
	model.CommonResponse
	CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponseModel
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponseModel).Reset()
}

// CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponseModel is 首公里揽收-快递预查询服务 成功返回结果
type CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_global_im_pickup_bigbag_express_prequery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// hsfResult
	HsfResult *HsfResult `json:"hsf_result,omitempty" xml:"hsf_result,omitempty"`
}

// Reset 清空结构体
func (m *CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.HsfResult = nil
}

var poolCainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse = sync.Pool{
	New: func() any {
		return new(CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse)
	},
}

// GetCainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse 从 sync.Pool 获取 CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse
func GetCainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse() *CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse {
	return poolCainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse.Get().(*CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse)
}

// ReleaseCainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse 将 CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse 保存到 sync.Pool
func ReleaseCainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse(v *CainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse) {
	v.Reset()
	poolCainiaoGlobalImPickupBigbagExpressPrequeryAPIResponse.Put(v)
}
