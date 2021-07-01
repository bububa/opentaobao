package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货入库单确认接口 API请求
taobao.qimen.returnorder.confirm

taobao.qimen.returnorder.confirm
*/
type TaobaoQimenReturnorderConfirmAPIRequest struct {
    model.Params
    // 
    _request   *ReturnOrderConfirmRequest
}

// 初始化TaobaoQimenReturnorderConfirmAPIRequest对象
func NewTaobaoQimenReturnorderConfirmRequest() *TaobaoQimenReturnorderConfirmAPIRequest{
    return &TaobaoQimenReturnorderConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnorderConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.returnorder.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnorderConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenReturnorderConfirmAPIRequest) SetRequest(_request *ReturnOrderConfirmRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenReturnorderConfirmAPIRequest) GetRequest() *ReturnOrderConfirmRequest {
    return r._request
}
