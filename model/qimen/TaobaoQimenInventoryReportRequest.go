package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存盘点通知接口 APIRequest
taobao.qimen.inventory.report

WMS调用奇门的接口,将库存盘点情况回传ERP
*/
type TaobaoQimenInventoryReportRequest struct {
    model.Params

    // 
    request   *InventoryReportRequest 

}

func NewTaobaoQimenInventoryReportRequest() *TaobaoQimenInventoryReportRequest{
    return &TaobaoQimenInventoryReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenInventoryReportRequest) GetApiMethodName() string {
    return "taobao.qimen.inventory.report"
}

func (r TaobaoQimenInventoryReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenInventoryReportRequest) SetRequest(request *InventoryReportRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenInventoryReportRequest) GetRequest() *InventoryReportRequest {
    return r.request
}

