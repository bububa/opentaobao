package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketSendAPIResponse 商家电子凭证发码成功回调接口 API返回值
// taobao.vmarket.eticket.send
//
// 外部商家成功发码回调接口
type TaobaoVmarketEticketSendAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketSendAPIResponseModel).Reset()
}

// TaobaoVmarketEticketSendAPIResponseModel is 商家电子凭证发码成功回调接口 成功返回结果
type TaobaoVmarketEticketSendAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_send_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0:失败；1:成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = 0
}

var poolTaobaoVmarketEticketSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketSendAPIResponse)
	},
}

// GetTaobaoVmarketEticketSendAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketSendAPIResponse
func GetTaobaoVmarketEticketSendAPIResponse() *TaobaoVmarketEticketSendAPIResponse {
	return poolTaobaoVmarketEticketSendAPIResponse.Get().(*TaobaoVmarketEticketSendAPIResponse)
}

// ReleaseTaobaoVmarketEticketSendAPIResponse 将 TaobaoVmarketEticketSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketSendAPIResponse(v *TaobaoVmarketEticketSendAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketSendAPIResponse.Put(v)
}
