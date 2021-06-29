package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
库存盘点通知接口 API请求
taobao.qimen.inventory.report

WMS调用奇门的接口,将库存盘点情况回传ERP
*/
type TaobaoQimenInventoryReportRequest struct {
    model.Params
    // 
    _request   *InventoryReportRequest
}

// 初始化TaobaoQimenInventoryReportRequest对象
func NewTaobaoQimenInventoryReportRequest() *TaobaoQimenInventoryReportRequest{
    return &TaobaoQimenInventoryReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenInventoryReportRequest) GetApiMethodName() string {
    return "taobao.qimen.inventory.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenInventoryReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenInventoryReportRequest) SetRequest(_request *InventoryReportRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenInventoryReportRequest) GetRequest() *InventoryReportRequest {
    return r._request
}
