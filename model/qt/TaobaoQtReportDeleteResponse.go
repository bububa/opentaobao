package qt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
质检报告删除接口 APIResponse
taobao.qt.report.delete

删除质检报告
*/
type TaobaoQtReportDeleteAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQtReportDeleteResponse `json:"qt_report_delete_response,omitempty"` 
    TaobaoQtReportDeleteResponse
}

/* model for simplify = false
type TaobaoQtReportDeleteResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoQtReportDeleteResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
