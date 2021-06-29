package fivee

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
国产商品发布 APIRequest
taobao.fivee.innerproduct.publish

资质共享平台国产商品发布
*/
type TaobaoFiveeInnerproductPublishRequest struct {
    model.Params

    // bu身份标识
    paramBucode   string 

    // 国产商品
    paramInnerProduct   *InnerProduct 

}

func NewTaobaoFiveeInnerproductPublishRequest() *TaobaoFiveeInnerproductPublishRequest{
    return &TaobaoFiveeInnerproductPublishRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoFiveeInnerproductPublishRequest) GetApiMethodName() string {
    return "taobao.fivee.innerproduct.publish"
}

func (r TaobaoFiveeInnerproductPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoFiveeInnerproductPublishRequest) SetParamBucode(paramBucode string) error {
    r.paramBucode = paramBucode
    r.Set("param_bucode", paramBucode)
    return nil
}

func (r TaobaoFiveeInnerproductPublishRequest) GetParamBucode() string {
    return r.paramBucode
}

func (r *TaobaoFiveeInnerproductPublishRequest) SetParamInnerProduct(paramInnerProduct *InnerProduct) error {
    r.paramInnerProduct = paramInnerProduct
    r.Set("param_inner_product", paramInnerProduct)
    return nil
}

func (r TaobaoFiveeInnerproductPublishRequest) GetParamInnerProduct() *InnerProduct {
    return r.paramInnerProduct
}

