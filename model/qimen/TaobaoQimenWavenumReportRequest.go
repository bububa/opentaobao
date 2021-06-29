package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单波次通知接口 API请求
taobao.qimen.wavenum.report

WMS调用奇门的接口,该接口用来给ERP或者OMS回传波次号及对应的发货单号，以支持商家货票同行、波次内包裹的处理等需求
*/
type TaobaoQimenWavenumReportRequest struct {
    model.Params
    // 
    _request   *WaveNumReportRequest
}

// 初始化TaobaoQimenWavenumReportRequest对象
func NewTaobaoQimenWavenumReportRequest() *TaobaoQimenWavenumReportRequest{
    return &TaobaoQimenWavenumReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenWavenumReportRequest) GetApiMethodName() string {
    return "taobao.qimen.wavenum.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenWavenumReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenWavenumReportRequest) SetRequest(_request *WaveNumReportRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenWavenumReportRequest) GetRequest() *WaveNumReportRequest {
    return r._request
}
