package ascp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse 快递送货上门能力同步/更新接口 API返回值
// taobao.logistics.express.delivery.send.ability.async
//
// 快递送货上门能力同步/更新接口
type TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse struct {
	model.CommonResponse
	TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponseModel).Reset()
}

// TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponseModel is 快递送货上门能力同步/更新接口 成功返回结果
type TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_delivery_send_ability_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	DeliverySendAbilityAsyncResponse *DeliverySendAblitiyAsyncResponse `json:"delivery_send_ability_async_response,omitempty" xml:"delivery_send_ability_async_response,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DeliverySendAbilityAsyncResponse = nil
}

var poolTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse)
	},
}

// GetTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse 从 sync.Pool 获取 TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse
func GetTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse() *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse {
	return poolTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse.Get().(*TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse)
}

// ReleaseTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse 将 TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse 保存到 sync.Pool
func ReleaseTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse(v *TaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse) {
	v.Reset()
	poolTaobaoLogisticsExpressDeliverySendAbilityAsyncAPIResponse.Put(v)
}
