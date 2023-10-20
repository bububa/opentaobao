package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketResendAPIResponse 外部合作商家重发电子凭证回调接口 API返回值
// taobao.vmarket.eticket.resend
//
// 外部合作商家重发电子凭证回调接口
type TaobaoVmarketEticketResendAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketResendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketResendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketResendAPIResponseModel).Reset()
}

// TaobaoVmarketEticketResendAPIResponseModel is 外部合作商家重发电子凭证回调接口 成功返回结果
type TaobaoVmarketEticketResendAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 0:失败，1:成功
	RetCode int64 `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketResendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = 0
}

var poolTaobaoVmarketEticketResendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketResendAPIResponse)
	},
}

// GetTaobaoVmarketEticketResendAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketResendAPIResponse
func GetTaobaoVmarketEticketResendAPIResponse() *TaobaoVmarketEticketResendAPIResponse {
	return poolTaobaoVmarketEticketResendAPIResponse.Get().(*TaobaoVmarketEticketResendAPIResponse)
}

// ReleaseTaobaoVmarketEticketResendAPIResponse 将 TaobaoVmarketEticketResendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketResendAPIResponse(v *TaobaoVmarketEticketResendAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketResendAPIResponse.Put(v)
}
