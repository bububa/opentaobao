package miniapp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCoinAwardDeliveryAPIResponse 淘金币奖励投放 API返回值
// taobao.coin.award.delivery
//
// 淘金币奖励投放
type TaobaoCoinAwardDeliveryAPIResponse struct {
	model.CommonResponse
	TaobaoCoinAwardDeliveryAPIResponseModel
}

// TaobaoCoinAwardDeliveryAPIResponseModel is 淘金币奖励投放 成功返回结果
type TaobaoCoinAwardDeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"coin_award_delivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 金币权益素材
	Result *TaobaoCoinAwardDeliveryResult `json:"result,omitempty" xml:"result,omitempty"`
}
