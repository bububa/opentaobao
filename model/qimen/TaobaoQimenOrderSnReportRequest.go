package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单SN通知接口 API请求
taobao.qimen.order.sn.report

WMS调用奇门的接口,在出库、发货、入库等场景下，ERP和WMS之间同步操作的SN列表
*/
type TaobaoQimenOrderSnReportRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenOrderSnReportRequest对象
func NewTaobaoQimenOrderSnReportRequest() *TaobaoQimenOrderSnReportRequest{
    return &TaobaoQimenOrderSnReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderSnReportRequest) GetApiMethodName() string {
    return "taobao.qimen.order.sn.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderSnReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderSnReportRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderSnReportRequest) GetRequest() *Request {
    return r.request
}
