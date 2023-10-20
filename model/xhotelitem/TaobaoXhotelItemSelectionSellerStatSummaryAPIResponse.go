package xhotelitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse 商家数据-选品整体概况 API返回值
// taobao.xhotel.item.selection.seller.stat.summary
//
// 商家数据-选品整体概况
type TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse struct {
	model.CommonResponse
	TaobaoXhotelItemSelectionSellerStatSummaryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoXhotelItemSelectionSellerStatSummaryAPIResponseModel).Reset()
}

// TaobaoXhotelItemSelectionSellerStatSummaryAPIResponseModel is 商家数据-选品整体概况 成功返回结果
type TaobaoXhotelItemSelectionSellerStatSummaryAPIResponseModel struct {
	XMLName xml.Name `xml:"xhotel_item_selection_seller_stat_summary_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *HsfResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoXhotelItemSelectionSellerStatSummaryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoXhotelItemSelectionSellerStatSummaryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse)
	},
}

// GetTaobaoXhotelItemSelectionSellerStatSummaryAPIResponse 从 sync.Pool 获取 TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse
func GetTaobaoXhotelItemSelectionSellerStatSummaryAPIResponse() *TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse {
	return poolTaobaoXhotelItemSelectionSellerStatSummaryAPIResponse.Get().(*TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse)
}

// ReleaseTaobaoXhotelItemSelectionSellerStatSummaryAPIResponse 将 TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoXhotelItemSelectionSellerStatSummaryAPIResponse(v *TaobaoXhotelItemSelectionSellerStatSummaryAPIResponse) {
	v.Reset()
	poolTaobaoXhotelItemSelectionSellerStatSummaryAPIResponse.Put(v)
}
