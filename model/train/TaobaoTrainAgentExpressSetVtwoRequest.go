package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息v2--增加鉴权校验 API请求
taobao.train.agent.express.set.vtwo

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetVtwoRequest struct {
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

// 初始化TaobaoTrainAgentExpressSetVtwoRequest对象
func NewTaobaoTrainAgentExpressSetVtwoRequest() *TaobaoTrainAgentExpressSetVtwoRequest{
    return &TaobaoTrainAgentExpressSetVtwoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.express.set.vtwo"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// MainOrderId Setter
// 订单号
func (r *TaobaoTrainAgentExpressSetVtwoRequest) SetMainOrderId(_mainOrderId int64) error {
    r._mainOrderId = _mainOrderId
    r.Set("main_order_id", _mainOrderId)
    return nil
}

// MainOrderId Getter
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetMainOrderId() int64 {
    return r._mainOrderId
}
// ExpressId Setter
// 物流单号
func (r *TaobaoTrainAgentExpressSetVtwoRequest) SetExpressId(_expressId string) error {
    r._expressId = _expressId
    r.Set("express_id", _expressId)
    return nil
}

// ExpressId Getter
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetExpressId() string {
    return r._expressId
}
// Addr Setter
// 发货地址
func (r *TaobaoTrainAgentExpressSetVtwoRequest) SetAddr(_addr string) error {
    r._addr = _addr
    r.Set("addr", _addr)
    return nil
}

// Addr Getter
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetAddr() string {
    return r._addr
}
// Mobile Setter
// 手机号
func (r *TaobaoTrainAgentExpressSetVtwoRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetMobile() string {
    return r._mobile
}
// AgentId Setter
// 代理商id
func (r *TaobaoTrainAgentExpressSetVtwoRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetAgentId() int64 {
    return r._agentId
}
// ExpressName Setter
// 物流公司:SF,EMS
func (r *TaobaoTrainAgentExpressSetVtwoRequest) SetExpressName(_expressName string) error {
    r._expressName = _expressName
    r.Set("express_name", _expressName)
    return nil
}

// ExpressName Getter
func (r TaobaoTrainAgentExpressSetVtwoRequest) GetExpressName() string {
    return r._expressName
}
