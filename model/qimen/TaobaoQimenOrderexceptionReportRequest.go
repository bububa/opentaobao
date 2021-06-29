package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单异常通知接口 API请求
taobao.qimen.orderexception.report

WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
*/
type TaobaoQimenOrderexceptionReportRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenOrderexceptionReportRequest对象
func NewTaobaoQimenOrderexceptionReportRequest() *TaobaoQimenOrderexceptionReportRequest{
    return &TaobaoQimenOrderexceptionReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderexceptionReportRequest) GetApiMethodName() string {
    return "taobao.qimen.orderexception.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderexceptionReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenOrderexceptionReportRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenOrderexceptionReportRequest) GetRequest() *Request {
    return r.request
}
