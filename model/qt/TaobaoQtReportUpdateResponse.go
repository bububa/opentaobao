package qt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新质检报告 APIResponse
taobao.qt.report.update

更新质检报告
*/
type TaobaoQtReportUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"qt_report_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"