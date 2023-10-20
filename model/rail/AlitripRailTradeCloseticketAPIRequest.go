package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriprailtradecloseticketAPIRequest 出票失败关单接口 API请求
// alitrip.rail.trade.closeticket
//
// 出票成功回调接口
type AlitriprailtradecloseticketAPIRequest struct {
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

// NewAlitriprailtradecloseticketRequest 初始化AlitriprailtradecloseticketAPIRequest对象
func NewAlitriprailtradecloseticketRequest() *AlitriprailtradecloseticketAPIRequest {
	return &AlitriprailtradecloseticketAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriprailtradecloseticketAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.trade.closeticket"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriprailtradecloseticketAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriprailtradecloseticketAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetErrorMsg is ErrorMsg Setter
// 出票失败原因
func (r *AlitriprailtradecloseticketAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// GetErrorMsg ErrorMsg Getter
func (r AlitriprailtradecloseticketAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// SetErrorCode is ErrorCode Setter
// 出票失败码
func (r *AlitriprailtradecloseticketAPIRequest) SetErrorCode(_errorCode string) error {
	r._errorCode = _errorCode
	r.Set("error_code", _errorCode)
	return nil
}

// GetErrorCode ErrorCode Getter
func (r AlitriprailtradecloseticketAPIRequest) GetErrorCode() string {
	return r._errorCode
}

// SetTpOrderId is TpOrderId Setter
// 平台订单号
func (r *AlitriprailtradecloseticketAPIRequest) SetTpOrderId(_tpOrderId int64) error {
	r._tpOrderId = _tpOrderId
	r.Set("tp_order_id", _tpOrderId)
	return nil
}

// GetTpOrderId TpOrderId Getter
func (r AlitriprailtradecloseticketAPIRequest) GetTpOrderId() int64 {
	return r._tpOrderId
}

// SetAgentId is AgentId Setter
// 代理商订单号
func (r *AlitriprailtradecloseticketAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitriprailtradecloseticketAPIRequest) GetAgentId() int64 {
	return r._agentId
}
