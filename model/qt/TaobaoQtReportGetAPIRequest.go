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
type TaobaoQtReportGetAPIRequest struct {
    model.Params
    // 质检编号
    _qtCode   string
}

// 初始化TaobaoQtReportGetAPIRequest对象
func NewTaobaoQtReportGetRequest() *TaobaoQtReportGetAPIRequest{
    return &TaobaoQtReportGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQtReportGetAPIRequest) GetApiMethodName() string {
    return "taobao.qt.report.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQtReportGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QtCode Setter
// 质检编号
func (r *TaobaoQtReportGetAPIRequest) SetQtCode(_qtCode string) error {
    r._qtCode = _qtCode
    r.Set("qt_code", _qtCode)
    return nil
}

// QtCode Getter
func (r TaobaoQtReportGetAPIRequest) GetQtCode() string {
    return r._qtCode
}
