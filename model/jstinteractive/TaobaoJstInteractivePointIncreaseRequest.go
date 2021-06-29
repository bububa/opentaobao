package jstinteractive

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
互动积分发放接口 APIRequest
taobao.jst.interactive.point.increase

向用户发放互动积分
*/
type TaobaoJstInteractivePointIncreaseRequest struct {
    model.Params

    // 发放的积分值
    amount   int64 

    // 幂等操作码，要确保唯一性，同一个操作码只能使用一次，避免重复操作
    operateCode   string 

}

func NewTaobaoJstInteractivePointIncreaseRequest() *TaobaoJstInteractivePointIncreaseRequest{
    return &TaobaoJstInteractivePointIncreaseRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstInteractivePointIncreaseRequest) GetApiMethodName() string {
    return "taobao.jst.interactive.point.increase"
}

func (r TaobaoJstInteractivePointIncreaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstInteractivePointIncreaseRequest) SetAmount(amount int64) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

func (r TaobaoJstInteractivePointIncreaseRequest) GetAmount() int64 {
    return r.amount
}

func (r *TaobaoJstInteractivePointIncreaseRequest) SetOperateCode(operateCode string) error {
    r.operateCode = operateCode
    r.Set("operate_code", operateCode)
    return nil
}

func (r TaobaoJstInteractivePointIncreaseRequest) GetOperateCode() string {
    return r.operateCode
}

