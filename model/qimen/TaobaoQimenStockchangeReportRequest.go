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
type TaobaoQimenStockchangeReportRequest struct {
    model.Params
    // 
    request   *StockChangeReportRequest
}

// 初始化TaobaoQimenStockchangeReportRequest对象
func NewTaobaoQimenStockchangeReportRequest() *TaobaoQimenStockchangeReportRequest{
    return &TaobaoQimenStockchangeReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenStockchangeReportRequest) GetApiMethodName() string {
    return "taobao.qimen.stockchange.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenStockchangeReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenStockchangeReportRequest) SetRequest(request *StockChangeReportRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenStockchangeReportRequest) GetRequest() *StockChangeReportRequest {
    return r.request
}
