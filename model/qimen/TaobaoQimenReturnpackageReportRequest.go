package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
退货包裹状态通知接口 APIRequest
taobao.qimen.returnpackage.report

退货包裹状态通知接口
*/
type TaobaoQimenReturnpackageReportRequest struct {
    model.Params

    // 
    request   *Request 

}

func NewTaobaoQimenReturnpackageReportRequest() *TaobaoQimenReturnpackageReportRequest{
    return &TaobaoQimenReturnpackageReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenReturnpackageReportRequest) GetApiMethodName() string {
    return "taobao.qimen.returnpackage.report"
}

func (r TaobaoQimenReturnpackageReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenReturnpackageReportRequest) SetRequest(request *Request) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenReturnpackageReportRequest) GetRequest() *Request {
    return r.request
}

