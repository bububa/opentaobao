package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailTradeCloseticketAPIRequest 出票失败关单接口 API请求
// alitrip.rail.trade.closeticket
//
// 出票成功回调接口
type AlitripRailTradeCloseticketAPIRequest struct {
	model.Params
	// 平台订单号
	_tpOrderId int64
	// 代理商订单号
	_agentId int64
	// 出票失败原因
	_errorMsg string
	// 出票失败码
	_errorCode string
}

// NewAlitripRailTradeCloseticketRequest 初始化AlitripRailTradeCloseticketAPIRequest对象
func NewAlitripRailTradeCloseticketRequest() *AlitripRailTradeCloseticketAPIRequest {
	return &AlitripRailTradeCloseticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailTradeCloseticketAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.trade.closeticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailTradeCloseticketAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TpOrderId Setter
// 平台订单号
func (r *AlitripRailTradeCloseticketAPIRequest) SetTpOrderId(_tpOrderId int64) error {
	r._tpOrderId = _tpOrderId
	r.Set("tp_order_id", _tpOrderId)
	return nil
}

// Get TpOrderId Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetTpOrderId() int64 {
	return r._tpOrderId
}

// Set is AgentId Setter
// 代理商订单号
func (r *AlitripRailTradeCloseticketAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// Get AgentId Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// Set is ErrorMsg Setter
// 出票失败原因
func (r *AlitripRailTradeCloseticketAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// Get ErrorMsg Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// Set is ErrorCode Setter
// 出票失败码
func (r *AlitripRailTradeCloseticketAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// Get ErrorCode Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetErrorCode() string {
	return r._errorCode
}
