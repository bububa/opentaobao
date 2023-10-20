package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentOrderSearchAPIResponse 【国际机票】订单列表查询 API返回值
// taobao.alitrip.ie.agent.order.search
//
// 根据指定条件查询国际机票订单列表
type TaobaoAlitripIeAgentOrderSearchAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripIeAgentOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripIeAgentOrderSearchAPIResponseModel).Reset()
}

// TaobaoAlitripIeAgentOrderSearchAPIResponseModel is 【国际机票】订单列表查询 成功返回结果
type TaobaoAlitripIeAgentOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_ie_agent_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 订单列表
	BaseOrderVos []IeBaseOrderVo `json:"base_order_vos,omitempty" xml:"base_order_vos>ie_base_order_vo,omitempty"`
	// 是否可以翻页查询
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
	// 请求成功标识
	QuerySuccess bool `json:"query_success,omitempty" xml:"query_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripIeAgentOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.BaseOrderVos = m.BaseOrderVos[:0]
	m.HasNext = false
	m.QuerySuccess = false
}

var poolTaobaoAlitripIeAgentOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripIeAgentOrderSearchAPIResponse)
	},
}

// GetTaobaoAlitripIeAgentOrderSearchAPIResponse 从 sync.Pool 获取 TaobaoAlitripIeAgentOrderSearchAPIResponse
func GetTaobaoAlitripIeAgentOrderSearchAPIResponse() *TaobaoAlitripIeAgentOrderSearchAPIResponse {
	return poolTaobaoAlitripIeAgentOrderSearchAPIResponse.Get().(*TaobaoAlitripIeAgentOrderSearchAPIResponse)
}

// ReleaseTaobaoAlitripIeAgentOrderSearchAPIResponse 将 TaobaoAlitripIeAgentOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripIeAgentOrderSearchAPIResponse(v *TaobaoAlitripIeAgentOrderSearchAPIResponse) {
	v.Reset()
	poolTaobaoAlitripIeAgentOrderSearchAPIResponse.Put(v)
}
