package traderate

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraderateFeedsGetAPIResponse 查询子订单对应的评价、追评以及语义标签 API返回值
// tmall.traderate.feeds.get
//
// 通过子订单ID获取天猫订单对应的评价，追评，以及对应的语义标签
type TmallTraderateFeedsGetAPIResponse struct {
	model.CommonResponse
	TmallTraderateFeedsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallTraderateFeedsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallTraderateFeedsGetAPIResponseModel).Reset()
}

// TmallTraderateFeedsGetAPIResponseModel is 查询子订单对应的评价、追评以及语义标签 成功返回结果
type TmallTraderateFeedsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_traderate_feeds_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回评价信息
	TmallRateInfo *TmallTraderateFeedsGetModel `json:"tmall_rate_info,omitempty" xml:"tmall_rate_info,omitempty"`
}

// Reset 清空结构体
func (m *TmallTraderateFeedsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TmallRateInfo = nil
}

var poolTmallTraderateFeedsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallTraderateFeedsGetAPIResponse)
	},
}

// GetTmallTraderateFeedsGetAPIResponse 从 sync.Pool 获取 TmallTraderateFeedsGetAPIResponse
func GetTmallTraderateFeedsGetAPIResponse() *TmallTraderateFeedsGetAPIResponse {
	return poolTmallTraderateFeedsGetAPIResponse.Get().(*TmallTraderateFeedsGetAPIResponse)
}

// ReleaseTmallTraderateFeedsGetAPIResponse 将 TmallTraderateFeedsGetAPIResponse 保存到 sync.Pool
func ReleaseTmallTraderateFeedsGetAPIResponse(v *TmallTraderateFeedsGetAPIResponse) {
	v.Reset()
	poolTmallTraderateFeedsGetAPIResponse.Put(v)
}
