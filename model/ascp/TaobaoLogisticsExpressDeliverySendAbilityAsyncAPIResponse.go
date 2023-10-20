package ascp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressdeliverysendabilityasyncAPIResponse 快递送货上门能力同步/更新接口 API返回值
// taobao.logistics.express.delivery.send.ability.async
//
// 快递送货上门能力同步/更新接口
type TaobaologisticsexpressdeliverysendabilityasyncAPIResponse struct {
	model.CommonResponse
	TaobaologisticsexpressdeliverysendabilityasyncAPIResponseModel
}

// TaobaologisticsexpressdeliverysendabilityasyncAPIResponseModel is 快递送货上门能力同步/更新接口 成功返回结果
type TaobaologisticsexpressdeliverysendabilityasyncAPIResponseModel struct {
	XMLName xml.Name `xml:"logistics_express_delivery_send_ability_async_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	DeliverySendAbilityAsyncResponse *DeliverySendAblitiyAsyncResponse `json:"delivery_send_ability_async_response,omitempty" xml:"delivery_send_ability_async_response,omitempty"`
}
