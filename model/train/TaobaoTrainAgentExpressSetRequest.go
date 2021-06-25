package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
线下票回填物流信息 APIRequest
taobao.train.agent.express.set

线下票回填物流信息服务
*/
type TaobaoTrainAgentExpressSetRequest struct {
    model.Params

    // 订单号
    mainOrderId   int64 

    // 物流单号
    expressId   string 

    // 发货地址
    addr   string 

    // 手机号
    mobile   string 

    // 代理商id
    agentId   int64 

    // 物流公司:SF,EMS
    expressName   string 

}

func NewTaobaoTrainAgentExpressSetRequest() *TaobaoTrainAgentExpressSetRequest{
    return &TaobaoTrainAgentExpressSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentExpressSetRequest) GetApiMethodName() string {
    return "taobao.train.agent.express.set"
}

func (r TaobaoTrainAgentExpressSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentExpressSetRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentExpressSetRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentExpressSetRequest) SetExpressId(expressId string) error {
    r.expressId = expressId
    r.Set("express_id", expressId)
    return nil
}

func (r TaobaoTrainAgentExpressSetRequest) GetExpressId() string {
    return r.expressId
}

func (r *TaobaoTrainAgentExpressSetRequest) SetAddr(addr string) error {
    r.addr = addr
    r.Set("addr", addr)
    return nil
}

func (r TaobaoTrainAgentExpressSetRequest) GetAddr() string {
    return r.addr
}

func (r *TaobaoTrainAgentExpressSetRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r TaobaoTrainAgentExpressSetRequest) GetMobile() string {
    return r.mobile
}

func (r *TaobaoTrainAgentExpressSetRequest) SetAgentId(agentId int64) error {
    r.agentId = agentId
    r.Set("agent_id", agentId)
    return nil
}

func (r TaobaoTrainAgentExpressSetRequest) GetAgentId() int64 {
    return r.agentId
}

func (r *TaobaoTrainAgentExpressSetRequest) SetExpressName(expressName string) error {
    r.expressName = expressName
    r.Set("express_name", expressName)
    return nil
}

func (r TaobaoTrainAgentExpressSetRequest) GetExpressName() string {
    return r.expressName
}

