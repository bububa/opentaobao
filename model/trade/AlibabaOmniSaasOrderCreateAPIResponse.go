package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOmniSaasOrderCreateAPIResponse 订单创建接口 API返回值
// alibaba.omni.saas.order.create
//
// 服务商利用现有的saas系统和阿里完成交易系统的对接
type AlibabaOmniSaasOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaOmniSaasOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOmniSaasOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOmniSaasOrderCreateAPIResponseModel).Reset()
}

// AlibabaOmniSaasOrderCreateAPIResponseModel is 订单创建接口 成功返回结果
type AlibabaOmniSaasOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_omni_saas_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// tradeNo
	TradeNo int64 `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// totalAmount
	TotalAmount int64 `json:"total_amount,omitempty" xml:"total_amount,omitempty"`
	// actualPayFee
	ActualPayFee int64 `json:"actual_pay_fee,omitempty" xml:"actual_pay_fee,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOmniSaasOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TradeNo = 0
	m.TotalAmount = 0
	m.ActualPayFee = 0
}

var poolAlibabaOmniSaasOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOmniSaasOrderCreateAPIResponse)
	},
}

// GetAlibabaOmniSaasOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaOmniSaasOrderCreateAPIResponse
func GetAlibabaOmniSaasOrderCreateAPIResponse() *AlibabaOmniSaasOrderCreateAPIResponse {
	return poolAlibabaOmniSaasOrderCreateAPIResponse.Get().(*AlibabaOmniSaasOrderCreateAPIResponse)
}

// ReleaseAlibabaOmniSaasOrderCreateAPIResponse 将 AlibabaOmniSaasOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOmniSaasOrderCreateAPIResponse(v *AlibabaOmniSaasOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaOmniSaasOrderCreateAPIResponse.Put(v)
}
