package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建订单 API请求
taobao.openmall.trade.create

创建Openmall订单
*/
type TaobaoOpenmallTradeCreateRequest struct {
    model.Params
    // 请求入参
    _paramTopTradeCreateDO   *TopTradeCreateDo
}

// 初始化TaobaoOpenmallTradeCreateRequest对象
func NewTaobaoOpenmallTradeCreateRequest() *TaobaoOpenmallTradeCreateRequest{
    return &TaobaoOpenmallTradeCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeCreateRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamTopTradeCreateDO Setter
// 请求入参
func (r *TaobaoOpenmallTradeCreateRequest) SetParamTopTradeCreateDO(_paramTopTradeCreateDO *TopTradeCreateDo) error {
    r._paramTopTradeCreateDO = _paramTopTradeCreateDO
    r.Set("param_top_trade_create_d_o", _paramTopTradeCreateDO)
    return nil
}

// ParamTopTradeCreateDO Getter
func (r TaobaoOpenmallTradeCreateRequest) GetParamTopTradeCreateDO() *TopTradeCreateDo {
    return r._paramTopTradeCreateDO
}
