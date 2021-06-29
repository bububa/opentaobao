package qt

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
上传质检报告 APIResponse
taobao.qt.report.add

上传质检报告
*/
type TaobaoQtReportAddAPIResponse struct {
    model.CommonResponse
    TaobaoQtReportAddResponse
}

type TaobaoQtReportAddResponse struct {
    XMLName xml.Name `xml:"qt_report_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 是否成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
}
