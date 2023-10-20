package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaReverseAPIResponse 电子凭证冲正接口 API返回值
// taobao.eticket.merchant.ma.reverse
//
// 电子凭证平台冲正接口
type TaobaoEticketMerchantMaReverseAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaReverseAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaReverseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantMaReverseAPIResponseModel).Reset()
}

// TaobaoEticketMerchantMaReverseAPIResponseModel is 电子凭证冲正接口 成功返回结果
type TaobaoEticketMerchantMaReverseAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_reverse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 回复结果
	RespBody *ReverseMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaReverseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantMaReverseAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantMaReverseAPIResponse)
	},
}

// GetTaobaoEticketMerchantMaReverseAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantMaReverseAPIResponse
func GetTaobaoEticketMerchantMaReverseAPIResponse() *TaobaoEticketMerchantMaReverseAPIResponse {
	return poolTaobaoEticketMerchantMaReverseAPIResponse.Get().(*TaobaoEticketMerchantMaReverseAPIResponse)
}

// ReleaseTaobaoEticketMerchantMaReverseAPIResponse 将 TaobaoEticketMerchantMaReverseAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantMaReverseAPIResponse(v *TaobaoEticketMerchantMaReverseAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantMaReverseAPIResponse.Put(v)
}
