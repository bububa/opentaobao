package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单缺货通知接口 API请求
taobao.qimen.itemlack.report

WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
*/
type TaobaoQimenItemlackReportRequest struct {
    model.Params
    // 
    _request   *ItemLackReportRequest
}

// 初始化TaobaoQimenItemlackReportRequest对象
func NewTaobaoQimenItemlackReportRequest() *TaobaoQimenItemlackReportRequest{
    return &TaobaoQimenItemlackReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenItemlackReportRequest) GetApiMethodName() string {
    return "taobao.qimen.itemlack.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenItemlackReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenItemlackReportRequest) SetRequest(_request *ItemLackReportRequest) error {
    r._request = _request
    r.Set("request", _request)
    return nil
}

// Request Getter
func (r TaobaoQimenItemlackReportRequest) GetRequest() *ItemLackReportRequest {
    return r._request
}
