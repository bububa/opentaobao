package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryAreaAPIResponse 地域报表查询 API返回值
// taobao.universalbp.report.query.area
//
// 地域报表查询
type TaobaoUniversalbpReportQueryAreaAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryAreaAPIResponseModel
}

// TaobaoUniversalbpReportQueryAreaAPIResponseModel is 地域报表查询 成功返回结果
type TaobaoUniversalbpReportQueryAreaAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_area_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryAreaTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
