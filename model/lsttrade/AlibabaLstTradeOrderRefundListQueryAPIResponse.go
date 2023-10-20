package lsttrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeOrderRefundListQueryAPIResponse 查询退款单列表(卖家视角) API返回值
// alibaba.lst.trade.order.refund.list.query
//
// 查询退款单列表(卖家视角)
type AlibabaLstTradeOrderRefundListQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeOrderRefundListQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeOrderRefundListQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeOrderRefundListQueryAPIResponseModel).Reset()
}

// AlibabaLstTradeOrderRefundListQueryAPIResponseModel is 查询退款单列表(卖家视角) 成功返回结果
type AlibabaLstTradeOrderRefundListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_order_refund_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 列表数据
	ContentList []Content `json:"content_list,omitempty" xml:"content_list>content,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 每页数量
	Size int64 `json:"size,omitempty" xml:"size,omitempty"`
	// 当前页
	Page int64 `json:"page,omitempty" xml:"page,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeOrderRefundListQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ContentList = m.ContentList[:0]
	m.Total = 0
	m.Size = 0
	m.Page = 0
}

var poolAlibabaLstTradeOrderRefundListQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeOrderRefundListQueryAPIResponse)
	},
}

// GetAlibabaLstTradeOrderRefundListQueryAPIResponse 从 sync.Pool 获取 AlibabaLstTradeOrderRefundListQueryAPIResponse
func GetAlibabaLstTradeOrderRefundListQueryAPIResponse() *AlibabaLstTradeOrderRefundListQueryAPIResponse {
	return poolAlibabaLstTradeOrderRefundListQueryAPIResponse.Get().(*AlibabaLstTradeOrderRefundListQueryAPIResponse)
}

// ReleaseAlibabaLstTradeOrderRefundListQueryAPIResponse 将 AlibabaLstTradeOrderRefundListQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeOrderRefundListQueryAPIResponse(v *AlibabaLstTradeOrderRefundListQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeOrderRefundListQueryAPIResponse.Put(v)
}
