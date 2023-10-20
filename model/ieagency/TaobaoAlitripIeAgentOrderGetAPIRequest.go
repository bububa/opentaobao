package ieagency

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripIeAgentOrderGetAPIRequest 【国际机票】查询订单详情 API请求
// taobao.alitrip.ie.agent.order.get
//
// 根据订单ID查询订单详情
type TaobaoAlitripIeAgentOrderGetAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 交易订单ID
	_tradeOrderId int64
}

// NewTaobaoAlitripIeAgentOrderGetRequest 初始化TaobaoAlitripIeAgentOrderGetAPIRequest对象
func NewTaobaoAlitripIeAgentOrderGetRequest() *TaobaoAlitripIeAgentOrderGetAPIRequest {
	return &TaobaoAlitripIeAgentOrderGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripIeAgentOrderGetAPIRequest) Reset() {
	r._agentId = 0
	r._tradeOrderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripIeAgentOrderGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.ie.agent.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripIeAgentOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripIeAgentOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *TaobaoAlitripIeAgentOrderGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoAlitripIeAgentOrderGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetTradeOrderId is TradeOrderId Setter
// 交易订单ID
func (r *TaobaoAlitripIeAgentOrderGetAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r TaobaoAlitripIeAgentOrderGetAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}

var poolTaobaoAlitripIeAgentOrderGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripIeAgentOrderGetRequest()
	},
}

// GetTaobaoAlitripIeAgentOrderGetRequest 从 sync.Pool 获取 TaobaoAlitripIeAgentOrderGetAPIRequest
func GetTaobaoAlitripIeAgentOrderGetAPIRequest() *TaobaoAlitripIeAgentOrderGetAPIRequest {
	return poolTaobaoAlitripIeAgentOrderGetAPIRequest.Get().(*TaobaoAlitripIeAgentOrderGetAPIRequest)
}

// ReleaseTaobaoAlitripIeAgentOrderGetAPIRequest 将 TaobaoAlitripIeAgentOrderGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripIeAgentOrderGetAPIRequest(v *TaobaoAlitripIeAgentOrderGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripIeAgentOrderGetAPIRequest.Put(v)
}
