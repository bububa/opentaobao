package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
仓内加工单确认接口 API请求
taobao.qimen.storeprocess.confirm

WMS调用奇门的接口,回传仓内加工单创建情况
*/
type TaobaoQimenStoreprocessConfirmRequest struct {
    model.Params
    // 
    _request   *StoreProcessConfirmRequest
}

// 初始化TaobaoQimenStoreprocessConfirmRequest对象
func NewTaobaoQimenStoreprocessConfirmRequest() *TaobaoQimenStoreprocessConfirmRequest{
    return &TaobaoQimenStoreprocessConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStoreprocessConfirmRequest) GetApiMethodName() string {
    return "taobao.qimen.storeprocess.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStoreprocessConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStoreprocessConfirmRequest) SetRequest(_request *StoreProcessConfirmRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenStoreprocessConfirmRequest) GetRequest() *StoreProcessConfirmRequest {
    return r._request
}
