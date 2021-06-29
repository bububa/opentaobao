package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单流水通知接口 API请求
taobao.qimen.orderprocess.report

WMS调用奇门的接口,将订单在仓库的状态回传给ERP；场景说明：仓库仓内操作状态回传给ERP, 比如打包操作完成时, 回传一个打 包完成的状态给到ERP, ERP自行决定如何处理。
*/
type TaobaoQimenOrderprocessReportRequest struct {
    model.Params
    // 
    request   *OrderProcessReportRequest
}

// 初始化TaobaoQimenOrderprocessReportRequest对象
func NewTaobaoQimenOrderprocessReportRequest() *TaobaoQimenOrderprocessReportRequest{
    return &TaobaoQimenOrderprocessReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessReportRequest) GetApiMethodName() string {
    return "taobao.qimen.orderprocess.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderprocessReportRequest) SetRequest(request *OrderProcessReportRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderprocessReportRequest) GetRequest() *OrderProcessReportRequest {
    return r.request
}
