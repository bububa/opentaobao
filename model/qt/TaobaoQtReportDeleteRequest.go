package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
质检报告删除接口 APIRequest
taobao.qt.report.delete

删除质检报告
*/
type TaobaoQtReportDeleteRequest struct {
    model.Params

    // 一个质检服务唯一标识质量检验单的编号
    qtCode   string 

}

func NewTaobaoQtReportDeleteRequest() *TaobaoQtReportDeleteRequest{
    return &TaobaoQtReportDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoQtReportDeleteRequest) GetApiMethodName() string {
    return "taobao.qt.report.delete"
}

func (r TaobaoQtReportDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoQtReportDeleteRequest) SetQtCode(qtCode string) error {
    r.qtCode = qtCode
    r.Set("qt_code", qtCode)
    return nil
}

func (r TaobaoQtReportDeleteRequest) GetQtCode() string {
    return r.qtCode
}

