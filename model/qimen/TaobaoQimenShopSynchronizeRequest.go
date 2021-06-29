package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
店铺同步接口 API请求
taobao.qimen.shop.synchronize

店铺同步接口描述
*/
type TaobaoQimenShopSynchronizeRequest struct {
    model.Params
    // 请求
    _request   *Request
}

// 初始化TaobaoQimenShopSynchronizeRequest对象
func NewTaobaoQimenShopSynchronizeRequest() *TaobaoQimenShopSynchronizeRequest{
    return &TaobaoQimenShopSynchronizeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenShopSynchronizeRequest) GetApiMethodName() string {
    return "taobao.qimen.shop.synchronize"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenShopSynchronizeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 请求
func (r *TaobaoQimenShopSynchronizeRequest) SetRequest(_request *Request) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenShopSynchronizeRequest) GetRequest() *Request {
    return r._request
}
