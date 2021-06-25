package qt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
上传质检报告 APIResponse
taobao.qt.report.add

上传质检报告
*/
type TaobaoQtReportAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoQtReportAddResponse `json:"taobao_qt_report_add_response,omitempty"`
}

type TaobaoQtReportAddResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
