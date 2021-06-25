package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单异常通知接口 APIRequest
taobao.qimen.orderexception.report

WMS调用奇门的接口,当WMS接收到ERP的发货指令时，由于种种原因（5.1.5说明了各种异常场景）可能无法完成发货。WMS通过调用此接口，通知ERP具体异常情况
*/
type TaobaoQimenOrderexceptionReportRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenOrderexceptionReportRequest() *TaobaoQimenOrderexceptionReportRequest{
    return &TaobaoQimenOrderexceptionReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenOrderexceptionReportRequest) GetApiMethodName() string {
    return "taobao.qimen.orderexception.report"
}

func (r TaobaoQimenOrderexceptionReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenOrderexceptionReportRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenOrderexceptionReportRequest) GetRequest() *Request {
    return r.request
}

