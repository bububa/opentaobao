package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单通知 APIRequest
taobao.qimen.transferorder.report

调拨单通知
*/
type TaobaoQimenTransferorderReportRequest struct {
    model.Params

    // 
    request   *TaobaoQimenTransferorderReportStruct 

}

func NewTaobaoQimenTransferorderReportRequest() *TaobaoQimenTransferorderReportRequest{
    return &TaobaoQimenTransferorderReportRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQimenTransferorderReportRequest) GetApiMethodName() string {
    return "taobao.qimen.transferorder.report"
}

func (r TaobaoQimenTransferorderReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQimenTransferorderReportRequest) SetRequest(request *TaobaoQimenTransferorderReportStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

func (r TaobaoQimenTransferorderReportRequest) GetRequest() *TaobaoQimenTransferorderReportStruct {
    return r.request
}

