package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
代理商出票中 APIRequest
taobao.train.agent.handleticket.confirm

代理商出票中
*/
type TaobaoTrainAgentHandleticketConfirmRequest struct {
    model.Params

    // 扩展参数
    extendParams   string 

    // 主站id
    mainOrderId   int64 

    // 代理商id
    sellerId   int64 

}

func NewTaobaoTrainAgentHandleticketConfirmRequest() *TaobaoTrainAgentHandleticketConfirmRequest{
    return &TaobaoTrainAgentHandleticketConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentHandleticketConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.handleticket.confirm"
}

func (r TaobaoTrainAgentHandleticketConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentHandleticketConfirmRequest) SetExtendParams(extendParams string) error {
    r.extendParams = extendParams
    r.Set("extend_params", extendParams)
    return nil
}

func (r TaobaoTrainAgentHandleticketConfirmRequest) GetExtendParams() string {
    return r.extendParams
}

func (r *TaobaoTrainAgentHandleticketConfirmRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoTrainAgentHandleticketConfirmRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

func (r *TaobaoTrainAgentHandleticketConfirmRequest) SetSellerId(sellerId int64) error {
    r.sellerId = sellerId
    r.Set("seller_id", sellerId)
    return nil
}

func (r TaobaoTrainAgentHandleticketConfirmRequest) GetSellerId() int64 {
    return r.sellerId
}

