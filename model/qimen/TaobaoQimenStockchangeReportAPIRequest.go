package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存异动通知接口 API请求
taobao.qimen.stockchange.report

WMS调用奇门的接口,将库存异动信息信息回传给ERP
*/
type TaobaoQimenStockchangeReportAPIRequest struct {
    model.Params
    // 
    _request   *StockChangeReportRequest
}

// 初始化TaobaoQimenStockchangeReportAPIRequest对象
func NewTaobaoQimenStockchangeReportRequest() *TaobaoQimenStockchangeReportAPIRequest{
    return &TaobaoQimenStockchangeReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockchangeReportAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.stockchange.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockchangeReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStockchangeReportAPIRequest) SetRequest(_request *StockChangeReportRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenStockchangeReportAPIRequest) GetRequest() *StockChangeReportRequest {
    return r._request
}
