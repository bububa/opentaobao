package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCrowdAPIResponse 人群报表查询 API返回值
// taobao.universalbp.report.query.crowd
//
// 人群报表查询
type TaobaoUniversalbpReportQueryCrowdAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryCrowdAPIResponseModel
}

// TaobaoUniversalbpReportQueryCrowdAPIResponseModel is 人群报表查询 成功返回结果
type TaobaoUniversalbpReportQueryCrowdAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_crowd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryCrowdTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
