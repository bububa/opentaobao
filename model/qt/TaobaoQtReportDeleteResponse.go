package qt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
质检报告删除接口 APIResponse
taobao.qt.report.delete

删除质检报告
*/
type TaobaoQtReportDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoQtReportDeleteResponse
}

type TaobaoQtReportDeleteResponse struct {
    XMLName xml.Name `xml:"qt_report_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
