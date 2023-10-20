package qt

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqtreportsgetAPIResponse 批量查询质检报告 API返回值
// taobao.qt.reports.get
//
// 批量查询质检报告，目前只支持查询qtType=11（天猫真假鉴定）类型的报告
type TaobaoqtreportsgetAPIResponse struct {
	model.CommonResponse
	TaobaoqtreportsgetAPIResponseModel
}

// TaobaoqtreportsgetAPIResponseModel is 批量查询质检报告 成功返回结果
type TaobaoqtreportsgetAPIResponseModel struct {
	XMLName xml.Name `xml:"qt_reports_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 质检报告列表
	Reports []QtReport `json:"reports,omitempty" xml:"reports>qt_report,omitempty"`
}
