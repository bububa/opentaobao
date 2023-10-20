package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTradeRefundInformAPIResponse 外部渠道通知淘鲜达退款成功接口 API返回值
// alibaba.wdk.trade.refund.inform
//
// 该接口用于外部渠道退款成功后，通知淘鲜达底层履约完成退款流程。
type AlibabaWdkTradeRefundInformAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTradeRefundInformAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTradeRefundInformAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTradeRefundInformAPIResponseModel).Reset()
}

// AlibabaWdkTradeRefundInformAPIResponseModel is 外部渠道通知淘鲜达退款成功接口 成功返回结果
type AlibabaWdkTradeRefundInformAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trade_refund_inform_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *InformRefundSuccessResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTradeRefundInformAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkTradeRefundInformAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTradeRefundInformAPIResponse)
	},
}

// GetAlibabaWdkTradeRefundInformAPIResponse 从 sync.Pool 获取 AlibabaWdkTradeRefundInformAPIResponse
func GetAlibabaWdkTradeRefundInformAPIResponse() *AlibabaWdkTradeRefundInformAPIResponse {
	return poolAlibabaWdkTradeRefundInformAPIResponse.Get().(*AlibabaWdkTradeRefundInformAPIResponse)
}

// ReleaseAlibabaWdkTradeRefundInformAPIResponse 将 AlibabaWdkTradeRefundInformAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTradeRefundInformAPIResponse(v *AlibabaWdkTradeRefundInformAPIResponse) {
	v.Reset()
	poolAlibabaWdkTradeRefundInformAPIResponse.Put(v)
}
