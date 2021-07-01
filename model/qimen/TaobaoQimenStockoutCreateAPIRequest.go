package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
出库单创建接口 API请求
taobao.qimen.stockout.create

ERP调用奇门接口，创建出库单信息
*/
type TaobaoQimenStockoutCreateAPIRequest struct {
    model.Params
    // 
    _request   *StockOutCreateRequest
}

// 初始化TaobaoQimenStockoutCreateAPIRequest对象
func NewTaobaoQimenStockoutCreateRequest() *TaobaoQimenStockoutCreateAPIRequest{
    return &TaobaoQimenStockoutCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockoutCreateAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.stockout.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockoutCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStockoutCreateAPIRequest) SetRequest(_request *StockOutCreateRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenStockoutCreateAPIRequest) GetRequest() *StockOutCreateRequest {
    return r._request
}
