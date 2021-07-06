package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentExpressSetVtwoAPIRequest 线下票回填物流信息v2--增加鉴权校验 API请求
// taobao.train.agent.express.set.vtwo
//
// 线下票回填物流信息服务
type TaobaoTrainAgentExpressSetVtwoAPIRequest struct {
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

// NewTaobaoTrainAgentExpressSetVtwoRequest 初始化TaobaoTrainAgentExpressSetVtwoAPIRequest对象
func NewTaobaoTrainAgentExpressSetVtwoRequest() *TaobaoTrainAgentExpressSetVtwoAPIRequest {
	return &TaobaoTrainAgentExpressSetVtwoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.express.set.vtwo"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExpressId is ExpressId Setter
// 物流单号
func (r *TaobaoTrainAgentExpressSetVtwoAPIRequest) SetExpressId(_expressId string) error {
	r._expressId = _expressId
	r.Set("express_id", _expressId)
	return nil
}

// GetExpressId ExpressId Getter
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetExpressId() string {
	return r._expressId
}

// SetAddr is Addr Setter
// 发货地址
func (r *TaobaoTrainAgentExpressSetVtwoAPIRequest) SetAddr(_addr string) error {
	r._addr = _addr
	r.Set("addr", _addr)
	return nil
}

// GetAddr Addr Getter
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetAddr() string {
	return r._addr
}

// SetMobile is Mobile Setter
// 手机号
func (r *TaobaoTrainAgentExpressSetVtwoAPIRequest) SetMobile(_mobile string) error {
	r._mobile = _mobile
	r.Set("mobile", _mobile)
	return nil
}

// GetMobile Mobile Getter
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetMobile() string {
	return r._mobile
}

// SetExpressName is ExpressName Setter
// 物流公司:SF,EMS
func (r *TaobaoTrainAgentExpressSetVtwoAPIRequest) SetExpressName(_expressName string) error {
	r._expressName = _expressName
	r.Set("express_name", _expressName)
	return nil
}

// GetExpressName ExpressName Getter
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetExpressName() string {
	return r._expressName
}

// SetMainOrderId is MainOrderId Setter
// 订单号
func (r *TaobaoTrainAgentExpressSetVtwoAPIRequest) SetMainOrderId(_mainOrderId int64) error {
	r._mainOrderId = _mainOrderId
	r.Set("main_order_id", _mainOrderId)
	return nil
}

// GetMainOrderId MainOrderId Getter
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetMainOrderId() int64 {
	return r._mainOrderId
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentExpressSetVtwoAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r TaobaoTrainAgentExpressSetVtwoAPIRequest) GetAgentId() int64 {
	return r._agentId
}
