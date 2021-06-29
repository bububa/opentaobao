package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营产品删除 API请求
taobao.alitrip.totoro.auxproduct.delete

廉航辅营产品删除接口
*/
type TaobaoAlitripTotoroAuxproductDeleteRequest struct {
    model.Params
    // 廉航辅营产品删除请求
    delAuxProductRq   *DelAuxProductRq
}

// 初始化TaobaoAlitripTotoroAuxproductDeleteRequest对象
func NewTaobaoAlitripTotoroAuxproductDeleteRequest() *TaobaoAlitripTotoroAuxproductDeleteRequest{
    return &TaobaoAlitripTotoroAuxproductDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoAlitripTotoroAuxproductDeleteRequest) GetApiMethodName() string {
    return "taobao.alitrip.totoro.auxproduct.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoAlitripTotoroAuxproductDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// DelAuxProductRq Setter
// 廉航辅营产品删除请求
func (r *TaobaoAlitripTotoroAuxproductDeleteRequest) SetDelAuxProductRq(delAuxProductRq *DelAuxProductRq) error {
    r.delAuxProductRq = delAuxProductRq
    r.Set("del_aux_product_rq", delAuxProductRq)
    return nil
}

// DelAuxProductRq Getter
func (r TaobaoAlitripTotoroAuxproductDeleteRequest) GetDelAuxProductRq() *DelAuxProductRq {
    return r.delAuxProductRq
}
