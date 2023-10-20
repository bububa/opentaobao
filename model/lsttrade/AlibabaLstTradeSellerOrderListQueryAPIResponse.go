package lsttrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstTradeSellerOrderListQueryAPIResponse 订单列表查看(卖家视角) API返回值
// alibaba.lst.trade.seller.order.list.query
//
// 卖家视角订单查询，查询授权经销商订单列表
type AlibabaLstTradeSellerOrderListQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstTradeSellerOrderListQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerOrderListQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstTradeSellerOrderListQueryAPIResponseModel).Reset()
}

// AlibabaLstTradeSellerOrderListQueryAPIResponseModel is 订单列表查看(卖家视角) 成功返回结果
type AlibabaLstTradeSellerOrderListQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_trade_seller_order_list_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstTradeSellerOrderListQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstTradeSellerOrderListQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstTradeSellerOrderListQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstTradeSellerOrderListQueryAPIResponse)
	},
}

// GetAlibabaLstTradeSellerOrderListQueryAPIResponse 从 sync.Pool 获取 AlibabaLstTradeSellerOrderListQueryAPIResponse
func GetAlibabaLstTradeSellerOrderListQueryAPIResponse() *AlibabaLstTradeSellerOrderListQueryAPIResponse {
	return poolAlibabaLstTradeSellerOrderListQueryAPIResponse.Get().(*AlibabaLstTradeSellerOrderListQueryAPIResponse)
}

// ReleaseAlibabaLstTradeSellerOrderListQueryAPIResponse 将 AlibabaLstTradeSellerOrderListQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstTradeSellerOrderListQueryAPIResponse(v *AlibabaLstTradeSellerOrderListQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstTradeSellerOrderListQueryAPIResponse.Put(v)
}
