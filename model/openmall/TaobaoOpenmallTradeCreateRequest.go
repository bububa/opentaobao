package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 APIRequest
taobao.openmall.trade.create

创建Openmall订单
*/
type TaobaoOpenmallTradeCreateRequest struct {
    model.Params

    // 请求入参
    paramTopTradeCreateDO   *TopTradeCreateDo 

}

func NewTaobaoOpenmallTradeCreateRequest() *TaobaoOpenmallTradeCreateRequest{
    return &TaobaoOpenmallTradeCreateRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallTradeCreateRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.create"
}

func (r TaobaoOpenmallTradeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallTradeCreateRequest) SetParamTopTradeCreateDO(paramTopTradeCreateDO *TopTradeCreateDo) error {
    r.paramTopTradeCreateDO = paramTopTradeCreateDO
    r.Set("param_top_trade_create_d_o", paramTopTradeCreateDO)
    return nil
}

func (r TaobaoOpenmallTradeCreateRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
    return r.paramTopTradeCreateDO
}

