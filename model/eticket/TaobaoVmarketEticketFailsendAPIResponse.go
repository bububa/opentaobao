package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketFailsendAPIResponse 无法发码回调 API返回值
// taobao.vmarket.eticket.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
type TaobaoVmarketEticketFailsendAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketFailsendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketFailsendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketFailsendAPIResponseModel).Reset()
}

// TaobaoVmarketEticketFailsendAPIResponseModel is 无法发码回调 成功返回结果
type TaobaoVmarketEticketFailsendAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_failsend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketFailsendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = 0
}

var poolTaobaoVmarketEticketFailsendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketFailsendAPIResponse)
	},
}

// GetTaobaoVmarketEticketFailsendAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketFailsendAPIResponse
func GetTaobaoVmarketEticketFailsendAPIResponse() *TaobaoVmarketEticketFailsendAPIResponse {
	return poolTaobaoVmarketEticketFailsendAPIResponse.Get().(*TaobaoVmarketEticketFailsendAPIResponse)
}

// ReleaseTaobaoVmarketEticketFailsendAPIResponse 将 TaobaoVmarketEticketFailsendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketFailsendAPIResponse(v *TaobaoVmarketEticketFailsendAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketFailsendAPIResponse.Put(v)
}
