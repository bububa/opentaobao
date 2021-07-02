package flight

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripTripvpAgentOrderGetAPIResponse 廉航辅营正向订单查询详情接口 API返回值
// alitrip.tripvp.agent.order.get
//
// 【国际机票】查询辅营订单详情
type AlitripTripvpAgentOrderGetAPIResponse struct {
	model.CommonResponse
	AlitripTripvpAgentOrderGetAPIResponseModel
}

// AlitripTripvpAgentOrderGetAPIResponseModel is 廉航辅营正向订单查询详情接口 成功返回结果
type AlitripTripvpAgentOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_tripvp_agent_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// orderVO
	OrderVo *VirProOrderVo `json:"order_vo,omitempty" xml:"order_vo,omitempty"`
	// pageSize
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
}
