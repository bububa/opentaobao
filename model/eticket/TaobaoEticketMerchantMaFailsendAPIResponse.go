package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaFailsendAPIResponse 码商发码失败回调接口 API返回值
// taobao.eticket.merchant.ma.failsend
//
// 针对一次发码通知，码商无法完成发码，则可以通过此接口告知电子凭证
type TaobaoEticketMerchantMaFailsendAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaFailsendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaFailsendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantMaFailsendAPIResponseModel).Reset()
}

// TaobaoEticketMerchantMaFailsendAPIResponseModel is 码商发码失败回调接口 成功返回结果
type TaobaoEticketMerchantMaFailsendAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_failsend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 回复参数
	RespBody *SendFailCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaFailsendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantMaFailsendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantMaFailsendAPIResponse)
	},
}

// GetTaobaoEticketMerchantMaFailsendAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantMaFailsendAPIResponse
func GetTaobaoEticketMerchantMaFailsendAPIResponse() *TaobaoEticketMerchantMaFailsendAPIResponse {
	return poolTaobaoEticketMerchantMaFailsendAPIResponse.Get().(*TaobaoEticketMerchantMaFailsendAPIResponse)
}

// ReleaseTaobaoEticketMerchantMaFailsendAPIResponse 将 TaobaoEticketMerchantMaFailsendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantMaFailsendAPIResponse(v *TaobaoEticketMerchantMaFailsendAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantMaFailsendAPIResponse.Put(v)
}
