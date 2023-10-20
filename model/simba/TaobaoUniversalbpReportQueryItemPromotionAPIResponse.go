package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryItemPromotionAPIResponse 宝贝主体报表查询 API返回值
// taobao.universalbp.report.query.item.promotion
//
// 宝贝主体报表查询
type TaobaoUniversalbpReportQueryItemPromotionAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryItemPromotionAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryItemPromotionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryItemPromotionAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryItemPromotionAPIResponseModel is 宝贝主体报表查询 成功返回结果
type TaobaoUniversalbpReportQueryItemPromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_item_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryItemPromotionTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryItemPromotionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryItemPromotionAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryItemPromotionAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryItemPromotionAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryItemPromotionAPIResponse
func GetTaobaoUniversalbpReportQueryItemPromotionAPIResponse() *TaobaoUniversalbpReportQueryItemPromotionAPIResponse {
	return poolTaobaoUniversalbpReportQueryItemPromotionAPIResponse.Get().(*TaobaoUniversalbpReportQueryItemPromotionAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryItemPromotionAPIResponse 将 TaobaoUniversalbpReportQueryItemPromotionAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryItemPromotionAPIResponse(v *TaobaoUniversalbpReportQueryItemPromotionAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryItemPromotionAPIResponse.Put(v)
}
