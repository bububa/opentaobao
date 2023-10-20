package lsttrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOrderDetailQueryAPIResponse 订单详情查看(卖家视角) API返回值
// alibaba.lst.trade.seller.order.detail.query
//
// 订单详情查看(卖家视角)
type AlibabaLstTradeSellerOrderDetailQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeSellerOrderDetailQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerOrderDetailQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeSellerOrderDetailQueryAPIResponseModel).Reset()
}

// AlibabaLstTradeSellerOrderDetailQueryAPIResponseModel is 订单详情查看(卖家视角) 成功返回结果
type AlibabaLstTradeSellerOrderDetailQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_seller_order_detail_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *AlibabaLstTradeSellerOrderDetailQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerOrderDetailQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstTradeSellerOrderDetailQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOrderDetailQueryAPIResponse)
	},
}

// GetAlibabaLstTradeSellerOrderDetailQueryAPIResponse 从 sync.Pool 获取 AlibabaLstTradeSellerOrderDetailQueryAPIResponse
func GetAlibabaLstTradeSellerOrderDetailQueryAPIResponse() *AlibabaLstTradeSellerOrderDetailQueryAPIResponse {
	return poolAlibabaLstTradeSellerOrderDetailQueryAPIResponse.Get().(*AlibabaLstTradeSellerOrderDetailQueryAPIResponse)
}

// ReleaseAlibabaLstTradeSellerOrderDetailQueryAPIResponse 将 AlibabaLstTradeSellerOrderDetailQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeSellerOrderDetailQueryAPIResponse(v *AlibabaLstTradeSellerOrderDetailQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeSellerOrderDetailQueryAPIResponse.Put(v)
}
