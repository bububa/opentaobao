package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建OC订单 API请求
taobao.oc.order.create

创建OC订单接口
*/
type TaobaoOcOrderCreateRequest struct {
    model.Params
    // OC订单
    _paramOCOrder   *OcOrder
}

// 初始化TaobaoOcOrderCreateRequest对象
func NewTaobaoOcOrderCreateRequest() *TaobaoOcOrderCreateRequest{
    return &TaobaoOcOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOcOrderCreateRequest) GetApiMethodName() string {
    return "taobao.oc.order.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOcOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamOCOrder Setter
// OC订单
func (r *TaobaoOcOrderCreateRequest) SetParamOCOrder(_paramOCOrder *OcOrder) error {
    r._paramOCOrder = _paramOCOrder
    r.Set("param_o_c_order", _paramOCOrder)
    return nil
}

// ParamOCOrder Getter
func (r TaobaoOcOrderCreateRequest) GetParamOCOrder() *OcOrder {
    return r._paramOCOrder
}
