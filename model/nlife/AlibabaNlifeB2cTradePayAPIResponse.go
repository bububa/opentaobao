package nlife

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaNlifeB2cTradePayAPIResponse 零售+平台支付订单 API返回值
// alibaba.nlife.b2c.trade.pay
//
// 零售+平台支付接口，外部商户调用此接口告知零售+支付结果，保持订单状态同步
type AlibabaNlifeB2cTradePayAPIResponse struct {
	model.CommonResponse
	AlibabaNlifeB2cTradePayAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradePayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaNlifeB2cTradePayAPIResponseModel).Reset()
}

// AlibabaNlifeB2cTradePayAPIResponseModel is 零售+平台支付订单 成功返回结果
type AlibabaNlifeB2cTradePayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_nlife_b2c_trade_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// gmtPayment
	GmtPayment string `json:"gmt_payment,omitempty" xml:"gmt_payment,omitempty"`
	// 扩展参数
	ExtendParams string `json:"extend_params,omitempty" xml:"extend_params,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaNlifeB2cTradePayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.GmtPayment = ""
	m.ExtendParams = ""
}

var poolAlibabaNlifeB2cTradePayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaNlifeB2cTradePayAPIResponse)
	},
}

// GetAlibabaNlifeB2cTradePayAPIResponse 从 sync.Pool 获取 AlibabaNlifeB2cTradePayAPIResponse
func GetAlibabaNlifeB2cTradePayAPIResponse() *AlibabaNlifeB2cTradePayAPIResponse {
	return poolAlibabaNlifeB2cTradePayAPIResponse.Get().(*AlibabaNlifeB2cTradePayAPIResponse)
}

// ReleaseAlibabaNlifeB2cTradePayAPIResponse 将 AlibabaNlifeB2cTradePayAPIResponse 保存到 sync.Pool
func ReleaseAlibabaNlifeB2cTradePayAPIResponse(v *AlibabaNlifeB2cTradePayAPIResponse) {
	v.Reset()
	poolAlibabaNlifeB2cTradePayAPIResponse.Put(v)
}
