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
type TaobaoQimenStockoutCreateRequest struct {
    model.Params
    // 
    _request   *StockOutCreateRequest
}

// 初始化TaobaoQimenStockoutCreateRequest对象
func NewTaobaoQimenStockoutCreateRequest() *TaobaoQimenStockoutCreateRequest{
    return &TaobaoQimenStockoutCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockoutCreateRequest) GetApiMethodName() string {
    return "taobao.qimen.stockout.create"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockoutCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStockoutCreateRequest) SetRequest(_request *StockOutCreateRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenStockoutCreateRequest) GetRequest() *StockOutCreateRequest {
    return r._request
}
