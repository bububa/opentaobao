package tmallchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelTradeRefundorderGetsAPIResponse 供应商查询退款单 API返回值
// tmall.channel.trade.refundorder.gets
//
// 供应商分页查询退款单
type TmallChannelTradeRefundorderGetsAPIResponse struct {
	model.CommonResponse
	TmallChannelTradeRefundorderGetsAPIResponseModel
}

// Reset 清空结构体
func (m *TmallChannelTradeRefundorderGetsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallChannelTradeRefundorderGetsAPIResponseModel).Reset()
}

// TmallChannelTradeRefundorderGetsAPIResponseModel is 供应商查询退款单 成功返回结果
type TmallChannelTradeRefundorderGetsAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_trade_refundorder_gets_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 包含的元素
	PageElements []TopChannelRefundDto `json:"page_elements,omitempty" xml:"page_elements>top_channel_refund_dto,omitempty"`
	// 每页显示数量
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 查询第几页
	PageNumber int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// 所有元素个数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TmallChannelTradeRefundorderGetsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PageElements = m.PageElements[:0]
	m.PageSize = 0
	m.PageNumber = 0
	m.TotalCount = 0
}

var poolTmallChannelTradeRefundorderGetsAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallChannelTradeRefundorderGetsAPIResponse)
	},
}

// GetTmallChannelTradeRefundorderGetsAPIResponse 从 sync.Pool 获取 TmallChannelTradeRefundorderGetsAPIResponse
func GetTmallChannelTradeRefundorderGetsAPIResponse() *TmallChannelTradeRefundorderGetsAPIResponse {
	return poolTmallChannelTradeRefundorderGetsAPIResponse.Get().(*TmallChannelTradeRefundorderGetsAPIResponse)
}

// ReleaseTmallChannelTradeRefundorderGetsAPIResponse 将 TmallChannelTradeRefundorderGetsAPIResponse 保存到 sync.Pool
func ReleaseTmallChannelTradeRefundorderGetsAPIResponse(v *TmallChannelTradeRefundorderGetsAPIResponse) {
	v.Reset()
	poolTmallChannelTradeRefundorderGetsAPIResponse.Put(v)
}
