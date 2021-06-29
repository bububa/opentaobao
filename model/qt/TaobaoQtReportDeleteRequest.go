package qt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
质检报告删除接口 API请求
taobao.qt.report.delete

删除质检报告
*/
type TaobaoQtReportDeleteRequest struct {
    model.Params
    // 一个质检服务唯一标识质量检验单的编号
    _qtCode   string
}

// 初始化TaobaoQtReportDeleteRequest对象
func NewTaobaoQtReportDeleteRequest() *TaobaoQtReportDeleteRequest{
    return &TaobaoQtReportDeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoQtReportDeleteRequest) GetApiMethodName() string {
    return "taobao.qt.report.delete"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoQtReportDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// QtCode Setter
// 一个质检服务唯一标识质量检验单的编号
func (r *TaobaoQtReportDeleteRequest) SetQtCode(_qtCode string) error {
    r._qtCode = _qtCode
    r.Set("qt_code", _qtCode)
    return nil
}

// QtCode Getter
func (r TaobaoQtReportDeleteRequest) GetQtCode() string {
    return r._qtCode
}
