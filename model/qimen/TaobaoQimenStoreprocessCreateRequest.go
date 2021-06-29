package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单创建接口 API请求
taobao.qimen.storeprocess.create

ERP调用奇门的接口,创建仓内加工单
*/
type TaobaoQimenStoreprocessCreateRequest struct {
    model.Params
    // 
    _request   *StoreProcessCreateRequest
}

// 初始化TaobaoQimenStoreprocessCreateRequest对象
func NewTaobaoQimenStoreprocessCreateRequest() *TaobaoQimenStoreprocessCreateRequest{
    return &TaobaoQimenStoreprocessCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreprocessCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.storeprocess.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreprocessCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStoreprocessCreateRequest) SetRequest(_request *StoreProcessCreateRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenStoreprocessCreateRequest) GetRequest() *StoreProcessCreateRequest {
    return r._request
}
