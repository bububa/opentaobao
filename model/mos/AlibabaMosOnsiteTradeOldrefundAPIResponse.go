package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMosOnsiteTradeOldrefundAPIResponse 线下新退款接口（专为老退款接口调用） API返回值
// alibaba.mos.onsite.trade.oldrefund
//
// 线下新退款接口（专为老退款接口调用）。新接口支付支付宝、手淘、天猫，老接口退款时，需要调用该接口退新单，并适配老接口响应参数返回
type AlibabaMosOnsiteTradeOldrefundAPIResponse struct {
	model.CommonResponse
	AlibabaMosOnsiteTradeOldrefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMosOnsiteTradeOldrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMosOnsiteTradeOldrefundAPIResponseModel).Reset()
}

// AlibabaMosOnsiteTradeOldrefundAPIResponseModel is 线下新退款接口（专为老退款接口调用） 成功返回结果
type AlibabaMosOnsiteTradeOldrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mos_onsite_trade_oldrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 交易退款响应
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMosOnsiteTradeOldrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaMosOnsiteTradeOldrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMosOnsiteTradeOldrefundAPIResponse)
	},
}

// GetAlibabaMosOnsiteTradeOldrefundAPIResponse 从 sync.Pool 获取 AlibabaMosOnsiteTradeOldrefundAPIResponse
func GetAlibabaMosOnsiteTradeOldrefundAPIResponse() *AlibabaMosOnsiteTradeOldrefundAPIResponse {
	return poolAlibabaMosOnsiteTradeOldrefundAPIResponse.Get().(*AlibabaMosOnsiteTradeOldrefundAPIResponse)
}

// ReleaseAlibabaMosOnsiteTradeOldrefundAPIResponse 将 AlibabaMosOnsiteTradeOldrefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMosOnsiteTradeOldrefundAPIResponse(v *AlibabaMosOnsiteTradeOldrefundAPIResponse) {
	v.Reset()
	poolAlibabaMosOnsiteTradeOldrefundAPIResponse.Put(v)
}
