package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotrainagentexpresssetAPIRequest 线下票回填物流信息 API请求
// taobao.train.agent.express.set
//
// 线下票回填物流信息服务
type TaobaotrainagentexpresssetAPIRequest struct {
	model.Params
	// 物流单号
	_expressId string
	// 发货地址
	_addr string
	// 手机号
	_mobile string
	// 物流公司:SF,EMS
	_expressName string
	// 订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// NewTaobaotrainagentexpresssetRequest 初始化TaobaotrainagentexpresssetAPIRequest对象
func NewTaobaotrainagentexpresssetRequest() *TaobaotrainagentexpresssetAPIRequest {
	return &TaobaotrainagentexpresssetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotrainagentexpresssetAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.express.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotrainagentexpresssetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotrainagentexpresssetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExpressId is ExpressId Setter
// 物流单号
func (r *TaobaotrainagentexpresssetAPIRequest) SetExpressId(_expressId string) error {
	r._expressId = _expressId
	r.Set("express_id", _expressId)
	return nil
}

// GetExpressId ExpressId Getter
func (r TaobaotrainagentexpresssetAPIRequest) GetExpressId() string {
	return r._expressId
}

// SetAddr is Addr Setter
// 发货地址
func (r *TaobaotrainagentexpresssetAPIRequest) SetAddr(_addr string) error {
	r._addr = _addr
	r.Set("addr", _addr)
	return nil
}

// GetAddr Addr Getter
func (r TaobaotrainagentexpresssetAPIRequest) GetAddr() string {
	return r._addr
}

// SetMobile is Mobile Setter
// 手机号
func (r *TaobaotrainagentexpresssetAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TaobaotrainagentexpresssetAPIRequest) GetMobile() string {
	return r._mobile
}

// SetExpressName is ExpressName Setter
// 物流公司:SF,EMS
func (r *TaobaotrainagentexpresssetAPIRequest) SetExpressName(_expressName string) error {
	r._expressName = _expressName
	r.Set("express_name", _expressName)
	return nil
}

// GetExpressName ExpressName Getter
func (r TaobaotrainagentexpresssetAPIRequest) GetExpressName() string {
	return r._expressName
}

// SetMainOrderId is MainOrderId Setter
// 订单号
func (r *TaobaotrainagentexpresssetAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaotrainagentexpresssetAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaotrainagentexpresssetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaotrainagentexpresssetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
