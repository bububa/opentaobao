package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse 其他主体报表查询 API返回值
// taobao.universalbp.report.query.not.item.promotion
//
// 其他主体报表查询
type TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryNotItemPromotionAPIResponseModel
}

// TaobaoUniversalbpReportQueryNotItemPromotionAPIResponseModel is 其他主体报表查询 成功返回结果
type TaobaoUniversalbpReportQueryNotItemPromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_not_item_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryNotItemPromotionTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
