package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentOrderGetAPIResponse 【国际机票】查询订单详情 API返回值
// taobao.alitrip.ie.agent.order.get
//
// 根据订单ID查询订单详情
type TaobaoAlitripIeAgentOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentOrderGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentOrderGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentOrderGetAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentOrderGetAPIResponseModel is 【国际机票】查询订单详情 成功返回结果
type TaobaoAlitripIeAgentOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据
	OrderVo *IeOrderVo `json:"order_vo,omitempty" xml:"order_vo,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentOrderGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderVo = nil
}

var poolTaobaoAlitripIeAgentOrderGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentOrderGetAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentOrderGetAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentOrderGetAPIResponse
func GetTaobaoAlitripIeAgentOrderGetAPIResponse() *TaobaoAlitripIeAgentOrderGetAPIResponse {
	return poolTaobaoAlitripIeAgentOrderGetAPIResponse.Get().(*TaobaoAlitripIeAgentOrderGetAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentOrderGetAPIResponse 将 TaobaoAlitripIeAgentOrderGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentOrderGetAPIResponse(v *TaobaoAlitripIeAgentOrderGetAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentOrderGetAPIResponse.Put(v)
}
