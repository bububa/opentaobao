package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单SN通知接口 APIRequest
taobao.qimen.sn.report

WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
*/
type TaobaoQimenSnReportRequest struct {
    model.Params

    // 
    request   *SnReportRequest 

}

func NewTaobaoQimenSnReportRequest() *TaobaoQimenSnReportRequest{
    return &TaobaoQimenSnReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenSnReportRequest) GetApiMethodName() string {
    return "taobao.qimen.sn.report"
}

func (r TaobaoQimenSnReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenSnReportRequest) SetRequest(request *SnReportRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenSnReportRequest) GetRequest() *SnReportRequest {
    return r.request
}

