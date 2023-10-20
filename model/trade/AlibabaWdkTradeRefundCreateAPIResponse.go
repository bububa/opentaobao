package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundCreateAPIResponse 外部渠道逆向订单创建 API返回值
// alibaba.wdk.trade.refund.create
//
// 该接口是创建退货订单的服务。当外部渠道发起退款后，调用此接口可以完成五道口底层交易、履约、配送等一系列流程进行退货。
type AlibabaWdkTradeRefundCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeRefundCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeRefundCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeRefundCreateAPIResponseModel).Reset()
}

// AlibabaWdkTradeRefundCreateAPIResponseModel is 外部渠道逆向订单创建 成功返回结果
type AlibabaWdkTradeRefundCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_refund_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RefundGoodsCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeRefundCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkTradeRefundCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeRefundCreateAPIResponse)
	},
}

// GetAlibabaWdkTradeRefundCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeRefundCreateAPIResponse
func GetAlibabaWdkTradeRefundCreateAPIResponse() *AlibabaWdkTradeRefundCreateAPIResponse {
	return poolAlibabaWdkTradeRefundCreateAPIResponse.Get().(*AlibabaWdkTradeRefundCreateAPIResponse)
}

// ReleaseAlibabaWdkTradeRefundCreateAPIResponse 将 AlibabaWdkTradeRefundCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeRefundCreateAPIResponse(v *AlibabaWdkTradeRefundCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeRefundCreateAPIResponse.Put(v)
}
