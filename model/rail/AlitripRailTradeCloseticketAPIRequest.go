package rail

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailTradeCloseticketAPIRequest 出票失败关单接口 API请求
// alitrip.rail.trade.closeticket
//
// 出票成功回调接口
type AlitripRailTradeCloseticketAPIRequest struct {
	model.Params
	// 出票失败原因
	_errorMsg string
	// 出票失败码
	_errorCode string
	// 平台订单号
	_tpOrderId int64
	// 代理商订单号
	_agentId int64
}

// NewAlitripRailTradeCloseticketRequest 初始化AlitripRailTradeCloseticketAPIRequest对象
func NewAlitripRailTradeCloseticketRequest() *AlitripRailTradeCloseticketAPIRequest {
	return &AlitripRailTradeCloseticketAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripRailTradeCloseticketAPIRequest) Reset() {
	r._errorMsg = ""
	r._errorCode = ""
	r._tpOrderId = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailTradeCloseticketAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.trade.closeticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailTradeCloseticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripRailTradeCloseticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErrorMsg is ErrorMsg Setter
// 出票失败原因
func (r *AlitripRailTradeCloseticketAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// GetErrorMsg ErrorMsg Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// SetErrorCode is ErrorCode Setter
// 出票失败码
func (r *AlitripRailTradeCloseticketAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetTpOrderId is TpOrderId Setter
// 平台订单号
func (r *AlitripRailTradeCloseticketAPIRequest) SetTpOrderId(_tpOrderId int64) error {
	r._tpOrderId = _tpOrderId
	r.Set("tp_order_id", _tpOrderId)
	return nil
}

// GetTpOrderId TpOrderId Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetTpOrderId() int64 {
	return r._tpOrderId
}

// SetAgentId is AgentId Setter
// 代理商订单号
func (r *AlitripRailTradeCloseticketAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripRailTradeCloseticketAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolAlitripRailTradeCloseticketAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripRailTradeCloseticketRequest()
	},
}

// GetAlitripRailTradeCloseticketRequest 从 sync.Pool 获取 AlitripRailTradeCloseticketAPIRequest
func GetAlitripRailTradeCloseticketAPIRequest() *AlitripRailTradeCloseticketAPIRequest {
	return poolAlitripRailTradeCloseticketAPIRequest.Get().(*AlitripRailTradeCloseticketAPIRequest)
}

// ReleaseAlitripRailTradeCloseticketAPIRequest 将 AlitripRailTradeCloseticketAPIRequest 放入 sync.Pool
func ReleaseAlitripRailTradeCloseticketAPIRequest(v *AlitripRailTradeCloseticketAPIRequest) {
	v.Reset()
	poolAlitripRailTradeCloseticketAPIRequest.Put(v)
}
