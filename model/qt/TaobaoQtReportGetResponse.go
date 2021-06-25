package qt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询质检报告 APIResponse
taobao.qt.report.get

质检报告查询
*/
type TaobaoQtReportGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQtReportGetResponse `json:"taobao_qt_report_get_response,omitempty"`
}

type TaobaoQtReportGetResponse struct {

    // 质检报告对象
    QtReport   *QtReport `json:"qt_report,omitempty"`

}
