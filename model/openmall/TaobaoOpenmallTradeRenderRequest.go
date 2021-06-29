package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渲染订单价格 APIRequest
taobao.openmall.trade.render

请求渲染订单价格
*/
type TaobaoOpenmallTradeRenderRequest struct {
    model.Params

    // 请求入参
    paramTopTradeCreateDO   *TopTradeCreateDo 

}

func NewTaobaoOpenmallTradeRenderRequest() *TaobaoOpenmallTradeRenderRequest{
    return &TaobaoOpenmallTradeRenderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenmallTradeRenderRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.render"
}

func (r TaobaoOpenmallTradeRenderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenmallTradeRenderRequest) SetParamTopTradeCreateDO(paramTopTradeCreateDO *TopTradeCreateDo) error {
    r.paramTopTradeCreateDO = paramTopTradeCreateDO
    r.Set("param_top_trade_create_d_o", paramTopTradeCreateDO)
    return nil
}

func (r TaobaoOpenmallTradeRenderRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
    return r.paramTopTradeCreateDO
}

