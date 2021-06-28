package qt

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新质检报告 APIResponse
taobao.qt.report.update

更新质检报告
*/
type TaobaoQtReportUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoQtReportUpdateResponse `json:"qt_report_update_response,omitempty"` 
    TaobaoQtReportUpdateResponse
}

/* model for simplify = false
type TaobaoQtReportUpdateResponse struct {

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoQtReportUpdateResponse struct {

    // 是否成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
