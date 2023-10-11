package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAdgroupAPIResponse 单元报表查询 API返回值
// taobao.universalbp.report.query.adgroup
//
// 单元报表查询
type TaobaoUniversalbpReportQueryAdgroupAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryAdgroupAPIResponseModel
}

// TaobaoUniversalbpReportQueryAdgroupAPIResponseModel is 单元报表查询 成功返回结果
type TaobaoUniversalbpReportQueryAdgroupAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_adgroup_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryAdgroupTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
