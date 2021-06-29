package qt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询质检报告 APIResponse
taobao.qt.report.get

质检报告查询
*/
type TaobaoQtReportGetAPIResponse struct {
    model.CommonResponse
    TaobaoQtReportGetResponse
}

type TaobaoQtReportGetResponse struct {
    XMLName xml.Name `xml:"qt_report_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 质检报告对象
    
    QtReport   *QtReport `json:"qt_report,omitempty" xml:"qt_report,omitempty"`

    
}
