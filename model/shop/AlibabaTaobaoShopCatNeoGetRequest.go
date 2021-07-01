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
type AlibabaTaobaoShopCatNeoGetAPIRequest struct {
    model.Params
    // 客户端鉴权虚拟api
    _unNamed   string
}

// 初始化AlibabaTaobaoShopCatNeoGetAPIRequest对象
func NewAlibabaTaobaoShopCatNeoGetRequest() *AlibabaTaobaoShopCatNeoGetAPIRequest{
    return &AlibabaTaobaoShopCatNeoGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaTaobaoShopCatNeoGetAPIRequest) GetApiMethodName() string {
    return "alibaba.taobao.shop.cat.neo.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaTaobaoShopCatNeoGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UnNamed Setter
// 客户端鉴权虚拟api
func (r *AlibabaTaobaoShopCatNeoGetAPIRequest) SetUnNamed(_unNamed string) error {
    r._unNamed = _unNamed
    r.Set("un_named", _unNamed)
    return nil
}

// UnNamed Getter
func (r AlibabaTaobaoShopCatNeoGetAPIRequest) GetUnNamed() string {
    return r._unNamed
}
