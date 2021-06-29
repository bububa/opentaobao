package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品投放 API请求
taobao.alitrip.totoro.auxproduct.push

廉航辅营产品投放接口
*/
type TaobaoAlitripTotoroAuxproductPushRequest struct {
    model.Params
    // 廉航辅营产品推送请求
    _pushAuxProductsRq   *PushAuxProductsRq
}

// 初始化TaobaoAlitripTotoroAuxproductPushRequest对象
func NewTaobaoAlitripTotoroAuxproductPushRequest() *TaobaoAlitripTotoroAuxproductPushRequest{
    return &TaobaoAlitripTotoroAuxproductPushRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTotoroAuxproductPushRequest) GetApiMethodName() string {
    return "taobao.alitrip.totoro.auxproduct.push"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTotoroAuxproductPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PushAuxProductsRq Setter
// 廉航辅营产品推送请求
func (r *TaobaoAlitripTotoroAuxproductPushRequest) SetPushAuxProductsRq(_pushAuxProductsRq *PushAuxProductsRq) error {
    r._pushAuxProductsRq = _pushAuxProductsRq
    r.Set("push_aux_products_rq", _pushAuxProductsRq)
    return nil
}

// PushAuxProductsRq Getter
func (r TaobaoAlitripTotoroAuxproductPushRequest) GetPushAuxProductsRq() *PushAuxProductsRq {
    return r._pushAuxProductsRq
}
