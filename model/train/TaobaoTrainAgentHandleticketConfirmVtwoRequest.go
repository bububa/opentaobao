package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商出票中v2--增加鉴权校验 APIRequest
taobao.train.agent.handleticket.confirm.vtwo

代理商出票中
*/
type TaobaoTrainAgentHandleticketConfirmVtwoRequest struct {
    model.Params

    // 扩展参数
    extendParams   string 

    // 主站id
    mainOrderId   int64 

    // 代理商id
    sellerId   int64 

}

func NewTaobaoTrainAgentHandleticketConfirmVtwoRequest() *TaobaoTrainAgentHandleticketConfirmVtwoRequest{
    return &TaobaoTrainAgentHandleticketConfirmVtwoRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetApiMethodName() string {
    return "taobao.train.agent.handleticket.confirm.vtwo"
}

func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentHandleticketConfirmVtwoRequest) SetExtendParams(extendParams string) error {
    r.extendParams = extendParams
    r.Set("extend_params", extendParams)
    return nil
}

func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetExtendParams() string {
    return r.extendParams
}

func (r *TaobaoTrainAgentHandleticketConfirmVtwoRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentHandleticketConfirmVtwoRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TaobaoTrainAgentHandleticketConfirmVtwoRequest) GetSellerId() int64 {
    return r.sellerId
}

