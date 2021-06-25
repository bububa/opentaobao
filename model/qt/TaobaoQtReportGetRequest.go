package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询质检报告 APIRequest
taobao.qt.report.get

质检报告查询
*/
type TaobaoQtReportGetRequest struct {
    model.Params

    // 质检编号
    qtCode   string 

}

func NewTaobaoQtReportGetRequest() *TaobaoQtReportGetRequest{
    return &TaobaoQtReportGetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQtReportGetRequest) GetApiMethodName() string {
    return "taobao.qt.report.get"
}

func (r TaobaoQtReportGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQtReportGetRequest) SetQtCode(qtCode string) error {
    r.qtCode = qtCode
    r.Set("qt_code", qtCode)
    return nil
}

func (r TaobaoQtReportGetRequest) GetQtCode() string {
    return r.qtCode
}

