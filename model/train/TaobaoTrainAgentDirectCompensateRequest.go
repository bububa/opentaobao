package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——订单关闭实际出票成功审计接口 APIRequest
taobao.train.agent.direct.compensate

代购直连订单平台关单但是代理商出票成功补偿接口
*/
type TaobaoTrainAgentDirectCompensateRequest struct {
    model.Params

    // 出票成功补偿入参
    compensateParam   *CompensateParam 

}

func NewTaobaoTrainAgentDirectCompensateRequest() *TaobaoTrainAgentDirectCompensateRequest{
    return &TaobaoTrainAgentDirectCompensateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentDirectCompensateRequest) GetApiMethodName() string {
    return "taobao.train.agent.direct.compensate"
}

func (r TaobaoTrainAgentDirectCompensateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentDirectCompensateRequest) SetCompensateParam(compensateParam *CompensateParam) error {
    r.compensateParam = compensateParam
    r.Set("compensate_param", compensateParam)
    return nil
}

func (r TaobaoTrainAgentDirectCompensateRequest) GetCompensateParam() *CompensateParam {
    return r.compensateParam
}

