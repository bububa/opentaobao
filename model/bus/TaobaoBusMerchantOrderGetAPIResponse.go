package bus

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBusMerchantOrderGetAPIResponse 商家侧查询订单详情 API返回值
// taobao.bus.merchant.order.get
//
// 商家侧查询订单详情
type TaobaoBusMerchantOrderGetAPIResponse struct {
	model.CommonResponse
	TaobaoBusMerchantOrderGetAPIResponseModel
}

// TaobaoBusMerchantOrderGetAPIResponseModel is 商家侧查询订单详情 成功返回结果
type TaobaoBusMerchantOrderGetAPIResponseModel struct {
	XMLName xml.Name `xml:"bus_merchant_order_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	ParamAgentQueryOrderRP *AgentQueryOrderRp `json:"param_agent_query_order_r_p,omitempty" xml:"param_agent_query_order_r_p,omitempty"`
}
