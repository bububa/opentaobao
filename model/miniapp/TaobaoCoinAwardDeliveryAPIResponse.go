package miniapp

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoCoinAwardDeliveryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCoinAwardDeliveryAPIResponseModel).Reset()
}

// TaobaoCoinAwardDeliveryAPIResponseModel is 淘金币奖励投放 成功返回结果
type TaobaoCoinAwardDeliveryAPIResponseModel struct {
	XMLName xml.Name `xml:"coin_award_delivery_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 金币权益素材
	Result *TaobaoCoinAwardDeliveryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCoinAwardDeliveryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoCoinAwardDeliveryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCoinAwardDeliveryAPIResponse)
	},
}

// GetTaobaoCoinAwardDeliveryAPIResponse 从 sync.Pool 获取 TaobaoCoinAwardDeliveryAPIResponse
func GetTaobaoCoinAwardDeliveryAPIResponse() *TaobaoCoinAwardDeliveryAPIResponse {
	return poolTaobaoCoinAwardDeliveryAPIResponse.Get().(*TaobaoCoinAwardDeliveryAPIResponse)
}

// ReleaseTaobaoCoinAwardDeliveryAPIResponse 将 TaobaoCoinAwardDeliveryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCoinAwardDeliveryAPIResponse(v *TaobaoCoinAwardDeliveryAPIResponse) {
	v.Reset()
	poolTaobaoCoinAwardDeliveryAPIResponse.Put(v)
}
