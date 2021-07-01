package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
渲染订单价格 API请求
taobao.openmall.trade.render

请求渲染订单价格
*/
type TaobaoOpenmallTradeRenderAPIRequest struct {
    model.Params
    // 请求入参
    _paramTopTradeCreateDO   *TopTradeCreateDo
}

// 初始化TaobaoOpenmallTradeRenderAPIRequest对象
func NewTaobaoOpenmallTradeRenderRequest() *TaobaoOpenmallTradeRenderAPIRequest{
    return &TaobaoOpenmallTradeRenderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeRenderAPIRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.render"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeRenderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopTradeCreateDO Setter
// 请求入参
func (r *TaobaoOpenmallTradeRenderAPIRequest) SetParamTopTradeCreateDO(_paramTopTradeCreateDO *TopTradeCreateDo) error {
    r._paramTopTradeCreateDO = _paramTopTradeCreateDO
    r.Set("param_top_trade_create_d_o", _paramTopTradeCreateDO)
    return nil
}

// ParamTopTradeCreateDO Getter
func (r TaobaoOpenmallTradeRenderAPIRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
    return r._paramTopTradeCreateDO
}
