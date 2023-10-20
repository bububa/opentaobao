package simba

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryNotItemPromotionAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryNotItemPromotionAPIResponseModel is 其他主体报表查询 成功返回结果
type TaobaoUniversalbpReportQueryNotItemPromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_not_item_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryNotItemPromotionTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryNotItemPromotionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryNotItemPromotionAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryNotItemPromotionAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse
func GetTaobaoUniversalbpReportQueryNotItemPromotionAPIResponse() *TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse {
	return poolTaobaoUniversalbpReportQueryNotItemPromotionAPIResponse.Get().(*TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryNotItemPromotionAPIResponse 将 TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryNotItemPromotionAPIResponse(v *TaobaoUniversalbpReportQueryNotItemPromotionAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryNotItemPromotionAPIResponse.Put(v)
}
