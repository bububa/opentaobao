package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundQueryAPIResponse 外部渠道查询退货订单详情接口 API返回值
// alibaba.wdk.trade.refund.query
//
// 该接口提供给外部渠道商家，比如欧尚外卖等查询退货订单详情，里面包含退货进度等信息。
type AlibabaWdkTradeRefundQueryAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeRefundQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeRefundQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeRefundQueryAPIResponseModel).Reset()
}

// AlibabaWdkTradeRefundQueryAPIResponseModel is 外部渠道查询退货订单详情接口 成功返回结果
type AlibabaWdkTradeRefundQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_refund_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	RefundGoodsQueryResult *RefundGoodsQueryResult `json:"refund_goods_query_result,omitempty" xml:"refund_goods_query_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeRefundQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RefundGoodsQueryResult = nil
}

var poolAlibabaWdkTradeRefundQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeRefundQueryAPIResponse)
	},
}

// GetAlibabaWdkTradeRefundQueryAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeRefundQueryAPIResponse
func GetAlibabaWdkTradeRefundQueryAPIResponse() *AlibabaWdkTradeRefundQueryAPIResponse {
	return poolAlibabaWdkTradeRefundQueryAPIResponse.Get().(*AlibabaWdkTradeRefundQueryAPIResponse)
}

// ReleaseAlibabaWdkTradeRefundQueryAPIResponse 将 AlibabaWdkTradeRefundQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeRefundQueryAPIResponse(v *AlibabaWdkTradeRefundQueryAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeRefundQueryAPIResponse.Put(v)
}
