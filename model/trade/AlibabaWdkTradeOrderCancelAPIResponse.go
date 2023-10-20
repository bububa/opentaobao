package trade

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdktradeordercancelAPIResponse 外部交易订单取消接口 API返回值
// alibaba.wdk.trade.order.cancel
//
// 通过该接口可以再盒马取消交易订单，并处理相关业务流程。主要用于和外部商户的订单进行同步和融合业务流程处理
type AlibabawdktradeordercancelAPIResponse struct {
	model.CommonResponse
	AlibabawdktradeordercancelAPIResponseModel
}

// AlibabawdktradeordercancelAPIResponseModel is 外部交易订单取消接口 成功返回结果
type AlibabawdktradeordercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_order_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Result *OrderResult `json:"result,omitempty" xml:"result,omitempty"`
}
