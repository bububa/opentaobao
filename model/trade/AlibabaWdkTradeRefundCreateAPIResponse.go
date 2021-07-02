package trade

import (
	"encoding/xml"

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

// AlibabaWdkTradeRefundCreateAPIResponseModel is 外部渠道逆向订单创建 成功返回结果
type AlibabaWdkTradeRefundCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_refund_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *RefundGoodsCreateResult `json:"result,omitempty" xml:"result,omitempty"`
}
