package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse 上门取退产能信息同步/更新 API返回值
// taobao.logistics.express.capacity.tms.async
//
// 上门取退产能信息同步/更新
type TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressCapacityTmsAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressCapacityTmsAsyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressCapacityTmsAsyncAPIResponseModel is 上门取退产能信息同步/更新 成功返回结果
type TaobaoLogisticsExpressCapacityTmsAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_capacity_tms_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	CapacityResponse *CapacityResponse `json:"capacity_response,omitempty" xml:"capacity_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressCapacityTmsAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CapacityResponse = nil
}

var poolTaobaoLogisticsExpressCapacityTmsAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressCapacityTmsAsyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse
func GetTaobaoLogisticsExpressCapacityTmsAsyncAPIResponse() *TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse {
	return poolTaobaoLogisticsExpressCapacityTmsAsyncAPIResponse.Get().(*TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressCapacityTmsAsyncAPIResponse 将 TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressCapacityTmsAsyncAPIResponse(v *TaobaoLogisticsExpressCapacityTmsAsyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressCapacityTmsAsyncAPIResponse.Put(v)
}
