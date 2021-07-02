package qt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQtReportGetAPIResponse 查询质检报告 API返回值
// taobao.qt.report.get
//
// 质检报告查询
type TaobaoQtReportGetAPIResponse struct {
	model.CommonResponse
	TaobaoQtReportGetAPIResponseModel
}

// TaobaoQtReportGetAPIResponseModel is 查询质检报告 成功返回结果
type TaobaoQtReportGetAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_report_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 质检报告对象
	QtReport *QtReport `json:"qt_report,omitempty" xml:"qt_report,omitempty"`
}
