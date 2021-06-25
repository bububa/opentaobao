package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品删除 APIRequest
taobao.alitrip.totoro.auxproduct.delete

廉航辅营产品删除接口
*/
type TaobaoAlitripTotoroAuxproductDeleteRequest struct {
    model.Params

    // 廉航辅营产品删除请求
    delAuxProductRq   *DelAuxProductRq 

}

func NewTaobaoAlitripTotoroAuxproductDeleteRequest() *TaobaoAlitripTotoroAuxproductDeleteRequest{
    return &TaobaoAlitripTotoroAuxproductDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoAlitripTotoroAuxproductDeleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.totoro.auxproduct.delete"
}

func (r TaobaoAlitripTotoroAuxproductDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoAlitripTotoroAuxproductDeleteRequest) SetDelAuxProductRq(delAuxProductRq *DelAuxProductRq) error {
    r.delAuxProductRq = delAuxProductRq
    r.Set("del_aux_product_rq", delAuxProductRq)
    return nil
}

func (r TaobaoAlitripTotoroAuxproductDeleteRequest) GetDelAuxProductRq() *DelAuxProductRq {
    return r.delAuxProductRq
}

