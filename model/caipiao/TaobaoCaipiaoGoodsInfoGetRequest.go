package caipiao

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据卖家id与appkey获取商品信息 APIRequest
taobao.caipiao.goods.info.get

根据卖家id与appkey获取商品信息。
*/
type TaobaoCaipiaoGoodsInfoGetRequest struct {
    model.Params

}

func NewTaobaoCaipiaoGoodsInfoGetRequest() *TaobaoCaipiaoGoodsInfoGetRequest{
    return &TaobaoCaipiaoGoodsInfoGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoCaipiaoGoodsInfoGetRequest) GetApiMethodName() string {
    return "taobao.caipiao.goods.info.get"
}

func (r TaobaoCaipiaoGoodsInfoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


