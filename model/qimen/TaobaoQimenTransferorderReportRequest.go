package qimen

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
调拨单通知 API请求
taobao.qimen.transferorder.report

调拨单通知
*/
type TaobaoQimenTransferorderReportRequest struct {
    model.Params
    // 
    request   *TaobaoQimenTransferorderReportStruct
}

// 初始化TaobaoQimenTransferorderReportRequest对象
func NewTaobaoQimenTransferorderReportRequest() *TaobaoQimenTransferorderReportRequest{
    return &TaobaoQimenTransferorderReportRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQimenTransferorderReportRequest) GetApiMethodName() string {
    return "taobao.qimen.transferorder.report"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQimenTransferorderReportRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Request Setter
// 
func (r *TaobaoQimenTransferorderReportRequest) SetRequest(request *TaobaoQimenTransferorderReportStruct) error {
    r.request = request
    r.Set("request", request)
    return nil
}

// Request Getter
func (r TaobaoQimenTransferorderReportRequest) GetRequest() *TaobaoQimenTransferorderReportStruct {
    return r.request
}
