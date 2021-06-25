package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品投放 APIRequest
taobao.alitrip.totoro.auxproduct.push

廉航辅营产品投放接口
*/
type TaobaoAlitripTotoroAuxproductPushRequest struct {
    model.Params

    // 廉航辅营产品推送请求
    pushAuxProductsRq   *PushAuxProductsRq 

}

func NewTaobaoAlitripTotoroAuxproductPushRequest() *TaobaoAlitripTotoroAuxproductPushRequest{
    return &TaobaoAlitripTotoroAuxproductPushRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTotoroAuxproductPushRequest) GetApiMethodName() string {
    return "taobao.alitrip.totoro.auxproduct.push"
}

func (r TaobaoAlitripTotoroAuxproductPushRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTotoroAuxproductPushRequest) SetPushAuxProductsRq(pushAuxProductsRq *PushAuxProductsRq) error {
    r.pushAuxProductsRq = pushAuxProductsRq
    r.Set("push_aux_products_rq", pushAuxProductsRq)
    return nil
}

func (r TaobaoAlitripTotoroAuxproductPushRequest) GetPushAuxProductsRq() *PushAuxProductsRq {
    return r.pushAuxProductsRq
}

