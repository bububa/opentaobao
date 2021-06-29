package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——订单关闭实际出票成功审计接口 API请求
taobao.train.agent.direct.compensate

代购直连订单平台关单但是代理商出票成功补偿接口
*/
type TaobaoTrainAgentDirectCompensateRequest struct {
    model.Params
    // 出票成功补偿入参
    compensateParam   *CompensateParam
}

// 初始化TaobaoTrainAgentDirectCompensateRequest对象
func NewTaobaoTrainAgentDirectCompensateRequest() *TaobaoTrainAgentDirectCompensateRequest{
    return &TaobaoTrainAgentDirectCompensateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentDirectCompensateRequest) GetApiMethodName() string {
    return "taobao.train.agent.direct.compensate"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentDirectCompensateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CompensateParam Setter
// 出票成功补偿入参
func (r *TaobaoTrainAgentDirectCompensateRequest) SetCompensateParam(compensateParam *CompensateParam) error {
    r.compensateParam = compensateParam
    r.Set("compensate_param", compensateParam)
    return nil
}

// CompensateParam Getter
func (r TaobaoTrainAgentDirectCompensateRequest) GetCompensateParam() *CompensateParam {
    return r.compensateParam
}
