package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
发货单缺货通知接口 APIRequest
taobao.qimen.itemlack.report

WMS调用奇门的接口,将商家在库某商品缺货的信息回传给ERP
*/
type TaobaoQimenItemlackReportRequest struct {
    model.Params

    // 
    request   *ItemLackReportRequest 

}

func NewTaobaoQimenItemlackReportRequest() *TaobaoQimenItemlackReportRequest{
    return &TaobaoQimenItemlackReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenItemlackReportRequest) GetApiMethodName() string {
    return "taobao.qimen.itemlack.report"
}

func (r TaobaoQimenItemlackReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenItemlackReportRequest) SetRequest(request *ItemLackReportRequest) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenItemlackReportRequest) GetRequest() *ItemLackReportRequest {
    return r.request
}

