package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCreativeAPIResponse 创意报表查询 API返回值
// taobao.universalbp.report.query.creative
//
// 创意报表查询
type TaobaoUniversalbpReportQueryCreativeAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryCreativeAPIResponseModel
}

// TaobaoUniversalbpReportQueryCreativeAPIResponseModel is 创意报表查询 成功返回结果
type TaobaoUniversalbpReportQueryCreativeAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_creative_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryCreativeTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
