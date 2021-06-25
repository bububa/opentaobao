package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存异动通知接口 APIRequest
taobao.qimen.stockchange.report

WMS调用奇门的接口,将库存异动信息信息回传给ERP
*/
type TaobaoQimenStockchangeReportRequest struct {
    model.Params

    // 
    request   *StockChangeReportRequest 

}

func NewTaobaoQimenStockchangeReportRequest() *TaobaoQimenStockchangeReportRequest{
    return &TaobaoQimenStockchangeReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenStockchangeReportRequest) GetApiMethodName() string {
    return "taobao.qimen.stockchange.report"
}

func (r TaobaoQimenStockchangeReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenStockchangeReportRequest) SetRequest(request *StockChangeReportRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenStockchangeReportRequest) GetRequest() *StockChangeReportRequest {
    return r.request
}

