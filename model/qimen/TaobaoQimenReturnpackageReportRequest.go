package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货包裹状态通知接口 API请求
taobao.qimen.returnpackage.report

退货包裹状态通知接口
*/
type TaobaoQimenReturnpackageReportRequest struct {
    model.Params
    // 
    request   *Request
}

// 初始化TaobaoQimenReturnpackageReportRequest对象
func NewTaobaoQimenReturnpackageReportRequest() *TaobaoQimenReturnpackageReportRequest{
    return &TaobaoQimenReturnpackageReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenReturnpackageReportRequest) GetApiMethodName() string {
    return "taobao.qimen.returnpackage.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenReturnpackageReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenReturnpackageReportRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenReturnpackageReportRequest) GetRequest() *Request {
    return r.request
}
