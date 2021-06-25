package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺mtop接口鉴权虚拟api APIRequest
alibaba.taobao.shop.cat.neo.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaTaobaoShopCatNeoGetRequest struct {
    model.Params

    // 客户端鉴权虚拟api
    unNamed   string 

}

func NewAlibabaTaobaoShopCatNeoGetRequest() *AlibabaTaobaoShopCatNeoGetRequest{
    return &AlibabaTaobaoShopCatNeoGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaTaobaoShopCatNeoGetRequest) GetApiMethodName() string {
    return "alibaba.taobao.shop.cat.neo.get"
}

func (r AlibabaTaobaoShopCatNeoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaTaobaoShopCatNeoGetRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

func (r AlibabaTaobaoShopCatNeoGetRequest) GetUnNamed() string {
    return r.unNamed
}

