package jipiao

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoJipiaoAgentOrderBdetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoJipiaoAgentOrderBdetailAPIResponseModel).Reset()
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

// Reset 清空结构体
func (m *TaobaoJipiaoAgentOrderBdetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Orders = m.Orders[:0]
	m.IsSuccess = false
}

var poolTaobaoJipiaoAgentOrderBdetailAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoJipiaoAgentOrderBdetailAPIResponse)
	},
}

// GetTaobaoJipiaoAgentOrderBdetailAPIResponse 从 sync.Pool 获取 TaobaoJipiaoAgentOrderBdetailAPIResponse
func GetTaobaoJipiaoAgentOrderBdetailAPIResponse() *TaobaoJipiaoAgentOrderBdetailAPIResponse {
	return poolTaobaoJipiaoAgentOrderBdetailAPIResponse.Get().(*TaobaoJipiaoAgentOrderBdetailAPIResponse)
}

// ReleaseTaobaoJipiaoAgentOrderBdetailAPIResponse 将 TaobaoJipiaoAgentOrderBdetailAPIResponse 保存到 sync.Pool
func ReleaseTaobaoJipiaoAgentOrderBdetailAPIResponse(v *TaobaoJipiaoAgentOrderBdetailAPIResponse) {
	v.Reset()
	poolTaobaoJipiaoAgentOrderBdetailAPIResponse.Put(v)
}
