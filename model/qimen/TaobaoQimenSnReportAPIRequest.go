package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单SN通知接口 API请求
taobao.qimen.sn.report

WMS调用奇门的接口,在仓库出库单后, 把SN信息回传给ERP
*/
type TaobaoQimenSnReportAPIRequest struct {
    model.Params
    // 
    _request   *SnReportRequest
}

// 初始化TaobaoQimenSnReportAPIRequest对象
func NewTaobaoQimenSnReportRequest() *TaobaoQimenSnReportAPIRequest{
    return &TaobaoQimenSnReportAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenSnReportAPIRequest) GetApiMethodName() string {
    return "taobao.qimen.sn.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenSnReportAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenSnReportAPIRequest) SetRequest(_request *SnReportRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenSnReportAPIRequest) GetRequest() *SnReportRequest {
    return r._request
}
