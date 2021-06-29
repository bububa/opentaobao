package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存预占取消接口 API请求
taobao.qimen.inventoryreserve.cancel

库存预占取消
*/
type TaobaoQimenInventoryreserveCancelRequest struct {
    model.Params
    // 
    _request   *Request
}

// 初始化TaobaoQimenInventoryreserveCancelRequest对象
func NewTaobaoQimenInventoryreserveCancelRequest() *TaobaoQimenInventoryreserveCancelRequest{
    return &TaobaoQimenInventoryreserveCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryreserveCancelRequest) GetApiMethodName() string {
    return "taobao.qimen.inventoryreserve.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryreserveCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenInventoryreserveCancelRequest) SetRequest(_request *Request) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenInventoryreserveCancelRequest) GetRequest() *Request {
    return r._request
}
