package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息 API请求
taobao.train.agent.express.set

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetRequest struct {
    model.Params
    // 订单号
    _mainOrderId   int64
    // 物流单号
    _expressId   string
    // 发货地址
    _addr   string
    // 手机号
    _mobile   string
    // 代理商id
    _agentId   int64
    // 物流公司:SF,EMS
    _expressName   string
}

// 初始化TaobaoTrainAgentExpressSetRequest对象
func NewTaobaoTrainAgentExpressSetRequest() *TaobaoTrainAgentExpressSetRequest{
    return &TaobaoTrainAgentExpressSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentExpressSetRequest) GetApiMethodName() string {
    return "taobao.train.agent.express.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentExpressSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 订单号
func (r *TaobaoTrainAgentExpressSetRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentExpressSetRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// ExpressId Setter
// 物流单号
func (r *TaobaoTrainAgentExpressSetRequest) SetExpressId(_expressId string) error {
    r._expressId = _expressId
    r.Set("express_id", _expressId)
    return nil
}

// ExpressId Getter
func (r TaobaoTrainAgentExpressSetRequest) GetExpressId() string {
    return r._expressId
}
// Addr Setter
// 发货地址
func (r *TaobaoTrainAgentExpressSetRequest) SetAddr(_addr string) error {
    r._addr = _addr
    r.Set("addr", _addr)
    return nil
}

// Addr Getter
func (r TaobaoTrainAgentExpressSetRequest) GetAddr() string {
    return r._addr
}
// Mobile Setter
// 手机号
func (r *TaobaoTrainAgentExpressSetRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoTrainAgentExpressSetRequest) GetMobile() string {
    return r._mobile
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentExpressSetRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentExpressSetRequest) GetAgentId() int64 {
    return r._agentId
}
// ExpressName Setter
// 物流公司:SF,EMS
func (r *TaobaoTrainAgentExpressSetRequest) SetExpressName(_expressName string) error {
    r._expressName = _expressName
    r.Set("express_name", _expressName)
    return nil
}

// ExpressName Getter
func (r TaobaoTrainAgentExpressSetRequest) GetExpressName() string {
    return r._expressName
}
