package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单波次通知接口 APIRequest
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
*/
type TaobaoQimenWavenumReportRequest struct {
    model.Params

    // 
    request   *WaveNumReportRequest 

}

func NewTaobaoQimenWavenumReportRequest() *TaobaoQimenWavenumReportRequest{
    return &TaobaoQimenWavenumReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenWavenumReportRequest) GetApiMethodName() string {
    return "taobao.qimen.wavenum.report"
}

func (r TaobaoQimenWavenumReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenWavenumReportRequest) SetRequest(request *WaveNumReportRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenWavenumReportRequest) GetRequest() *WaveNumReportRequest {
    return r.request
}

