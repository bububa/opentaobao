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
    request   *Request
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
func (r *TaobaoQimenShopSynchronizeRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenShopSynchronizeRequest) GetRequest() *Request {
    return r.request
}
