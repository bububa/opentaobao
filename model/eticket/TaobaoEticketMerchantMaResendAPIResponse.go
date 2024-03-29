package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaResendAPIResponse 电子凭证重发回调接口 API返回值
// taobao.eticket.merchant.ma.resend
//
// 码商重发电子凭证回调接口
type TaobaoEticketMerchantMaResendAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaResendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaResendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantMaResendAPIResponseModel).Reset()
}

// TaobaoEticketMerchantMaResendAPIResponseModel is 电子凭证重发回调接口 成功返回结果
type TaobaoEticketMerchantMaResendAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_resend_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 回复参数
	RespBody *SendMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaResendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantMaResendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantMaResendAPIResponse)
	},
}

// GetTaobaoEticketMerchantMaResendAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantMaResendAPIResponse
func GetTaobaoEticketMerchantMaResendAPIResponse() *TaobaoEticketMerchantMaResendAPIResponse {
	return poolTaobaoEticketMerchantMaResendAPIResponse.Get().(*TaobaoEticketMerchantMaResendAPIResponse)
}

// ReleaseTaobaoEticketMerchantMaResendAPIResponse 将 TaobaoEticketMerchantMaResendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantMaResendAPIResponse(v *TaobaoEticketMerchantMaResendAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantMaResendAPIResponse.Put(v)
}
