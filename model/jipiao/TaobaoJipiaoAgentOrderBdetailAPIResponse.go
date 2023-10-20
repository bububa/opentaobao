package jipiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJipiaoAgentOrderBdetailAPIResponse 【机票代理商订单】采购订单详情 API返回值
// taobao.jipiao.agent.order.bdetail
//
// 根据淘宝系统订单号获取订单详情信息
type TaobaoJipiaoAgentOrderBdetailAPIResponse struct {
	model.CommonResponse
	TaobaoJipiaoAgentOrderBdetailAPIResponseModel
}

// TaobaoJipiaoAgentOrderBdetailAPIResponseModel is 【机票代理商订单】采购订单详情 成功返回结果
type TaobaoJipiaoAgentOrderBdetailAPIResponseModel struct {
	XMLName xml.Name `xml:"jipiao_agent_order_bdetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 机票订单的详情列表，当前支持返回一个订单
	Orders []TripOrder `json:"orders,omitempty" xml:"orders>trip_order,omitempty"`
	// 返回操作成功失败信息
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
