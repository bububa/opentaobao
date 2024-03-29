package ieagency

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTripvpAgentOrderSearchAPIResponse 【国际机票】查询辅营订单列表 API返回值
// alitrip.tripvp.agent.order.search
//
// 【国际机票】查询辅营订单列表
type AlitripTripvpAgentOrderSearchAPIResponse struct {
	model.CommonResponse
	AlitripTripvpAgentOrderSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripTripvpAgentOrderSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripTripvpAgentOrderSearchAPIResponseModel).Reset()
}

// AlitripTripvpAgentOrderSearchAPIResponseModel is 【国际机票】查询辅营订单列表 成功返回结果
type AlitripTripvpAgentOrderSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tripvp_agent_order_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 辅营订单列表
	OrderVos []VirProOrderVo `json:"order_vos,omitempty" xml:"order_vos>vir_pro_order_vo,omitempty"`
	// 当前页码CurrentPage
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 是否有下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *AlitripTripvpAgentOrderSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.OrderVos = m.OrderVos[:0]
	m.CurrentPage = 0
	m.HasNext = false
}

var poolAlitripTripvpAgentOrderSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripTripvpAgentOrderSearchAPIResponse)
	},
}

// GetAlitripTripvpAgentOrderSearchAPIResponse 从 sync.Pool 获取 AlitripTripvpAgentOrderSearchAPIResponse
func GetAlitripTripvpAgentOrderSearchAPIResponse() *AlitripTripvpAgentOrderSearchAPIResponse {
	return poolAlitripTripvpAgentOrderSearchAPIResponse.Get().(*AlitripTripvpAgentOrderSearchAPIResponse)
}

// ReleaseAlitripTripvpAgentOrderSearchAPIResponse 将 AlitripTripvpAgentOrderSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripTripvpAgentOrderSearchAPIResponse(v *AlitripTripvpAgentOrderSearchAPIResponse) {
	v.Reset()
	poolAlitripTripvpAgentOrderSearchAPIResponse.Put(v)
}
