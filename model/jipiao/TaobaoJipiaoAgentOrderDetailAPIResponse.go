package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJipiaoAgentOrderDetailAPIResponse
【机票代理商订单】订单详情 API返回值
taobao.jipiao.agent.order.detail

根据淘宝系统订单号获取订单详情信息 */
type TaobaoJipiaoAgentOrderDetailAPIResponse struct {
	model.CommonResponse
	TaobaoJipiaoAgentOrderDetailAPIResponseModel
}

// TaobaoJipiaoAgentOrderDetailAPIResponseModel is 【机票代理商订单】订单详情 成功返回结果
type TaobaoJipiaoAgentOrderDetailAPIResponseModel struct {
	XMLName xml.Name `xml:"jipiao_agent_order_detail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回操作成功失败信息
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
	// 机票订单的详情列表，当前支持返回一个订单
	Orders []TripOrder `json:"orders,omitempty" xml:"orders>trip_order,omitempty"`
}
