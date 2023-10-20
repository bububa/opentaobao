package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaSendAPIResponse 码商发码成功回调接口 API返回值
// taobao.eticket.merchant.ma.send
//
// 码商发码成功回调接口
type TaobaoEticketMerchantMaSendAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaSendAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaSendAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantMaSendAPIResponseModel).Reset()
}

// TaobaoEticketMerchantMaSendAPIResponseModel is 码商发码成功回调接口 成功返回结果
type TaobaoEticketMerchantMaSendAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_send_response"`
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
func (m *TaobaoEticketMerchantMaSendAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantMaSendAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantMaSendAPIResponse)
	},
}

// GetTaobaoEticketMerchantMaSendAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantMaSendAPIResponse
func GetTaobaoEticketMerchantMaSendAPIResponse() *TaobaoEticketMerchantMaSendAPIResponse {
	return poolTaobaoEticketMerchantMaSendAPIResponse.Get().(*TaobaoEticketMerchantMaSendAPIResponse)
}

// ReleaseTaobaoEticketMerchantMaSendAPIResponse 将 TaobaoEticketMerchantMaSendAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantMaSendAPIResponse(v *TaobaoEticketMerchantMaSendAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantMaSendAPIResponse.Put(v)
}
