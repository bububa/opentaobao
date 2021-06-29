package shop

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺mtop接口鉴权虚拟api API请求
alibaba.taobao.shop.cat.neo.get

获取优惠券信息，仅作客户端鉴权虚拟api使用
*/
type AlibabaTaobaoShopCatNeoGetRequest struct {
    model.Params
    // 客户端鉴权虚拟api
    unNamed   string
}

// 初始化AlibabaTaobaoShopCatNeoGetRequest对象
func NewAlibabaTaobaoShopCatNeoGetRequest() *AlibabaTaobaoShopCatNeoGetRequest{
    return &AlibabaTaobaoShopCatNeoGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoShopCatNeoGetRequest) GetApiMethodName() string {
    return "alibaba.taobao.shop.cat.neo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoShopCatNeoGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权虚拟api
func (r *AlibabaTaobaoShopCatNeoGetRequest) SetUnNamed(unNamed string) error {
    r.unNamed = unNamed
    r.Set("un_named", unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaTaobaoShopCatNeoGetRequest) GetUnNamed() string {
    return r.unNamed
}
