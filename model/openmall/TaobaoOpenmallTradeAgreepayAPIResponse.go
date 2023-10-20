package openmall

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeAgreepayAPIResponse openmall订单支付 API返回值
// taobao.openmall.trade.agreepay
//
// openmall订单支付
type TaobaoOpenmallTradeAgreepayAPIResponse struct {
	model.CommonResponse
	TaobaoOpenmallTradeAgreepayAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeAgreepayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpenmallTradeAgreepayAPIResponseModel).Reset()
}

// TaobaoOpenmallTradeAgreepayAPIResponseModel is openmall订单支付 成功返回结果
type TaobaoOpenmallTradeAgreepayAPIResponseModel struct {
	XMLName xml.Name `xml:"openmall_trade_agreepay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpenmallTradeAgreepayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolTaobaoOpenmallTradeAgreepayAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpenmallTradeAgreepayAPIResponse)
	},
}

// GetTaobaoOpenmallTradeAgreepayAPIResponse 从 sync.Pool 获取 TaobaoOpenmallTradeAgreepayAPIResponse
func GetTaobaoOpenmallTradeAgreepayAPIResponse() *TaobaoOpenmallTradeAgreepayAPIResponse {
	return poolTaobaoOpenmallTradeAgreepayAPIResponse.Get().(*TaobaoOpenmallTradeAgreepayAPIResponse)
}

// ReleaseTaobaoOpenmallTradeAgreepayAPIResponse 将 TaobaoOpenmallTradeAgreepayAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpenmallTradeAgreepayAPIResponse(v *TaobaoOpenmallTradeAgreepayAPIResponse) {
	v.Reset()
	poolTaobaoOpenmallTradeAgreepayAPIResponse.Put(v)
}
