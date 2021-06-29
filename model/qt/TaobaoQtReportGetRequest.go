package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询质检报告 API请求
taobao.qt.report.get

质检报告查询
*/
type TaobaoQtReportGetRequest struct {
    model.Params
    // 质检编号
    qtCode   string
}

// 初始化TaobaoQtReportGetRequest对象
func NewTaobaoQtReportGetRequest() *TaobaoQtReportGetRequest{
    return &TaobaoQtReportGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQtReportGetRequest) GetApiMethodName() string {
    return "taobao.qt.report.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQtReportGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QtCode Setter
// 质检编号
func (r *TaobaoQtReportGetRequest) SetQtCode(qtCode string) error {
    r.qtCode = qtCode
    r.Set("qt_code", qtCode)
    return nil
}

// QtCode Getter
func (r TaobaoQtReportGetRequest) GetQtCode() string {
    return r.qtCode
}
