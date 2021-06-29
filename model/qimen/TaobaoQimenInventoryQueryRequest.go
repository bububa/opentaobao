package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存查询接口（多商品） API请求
taobao.qimen.inventory.query

ERP调用奇门的接口,查询商品的库存量
*/
type TaobaoQimenInventoryQueryRequest struct {
    model.Params
    // 
    _request   *InventoryQueryRequest
}

// 初始化TaobaoQimenInventoryQueryRequest对象
func NewTaobaoQimenInventoryQueryRequest() *TaobaoQimenInventoryQueryRequest{
    return &TaobaoQimenInventoryQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryQueryRequest) GetApiMethodName() string {
    return "taobao.qimen.inventory.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenInventoryQueryRequest) SetRequest(_request *InventoryQueryRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenInventoryQueryRequest) GetRequest() *InventoryQueryRequest {
    return r._request
}
