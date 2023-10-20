package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaDelayAPIResponse 凭证延期 API返回值
// taobao.eticket.merchant.ma.delay
//
// 订单延期
type TaobaoEticketMerchantMaDelayAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaDelayAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaDelayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantMaDelayAPIResponseModel).Reset()
}

// TaobaoEticketMerchantMaDelayAPIResponseModel is 凭证延期 成功返回结果
type TaobaoEticketMerchantMaDelayAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_delay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误消息
	ResMsg string `json:"res_msg,omitempty" xml:"res_msg,omitempty"`
	// 错误码
	ResCode int64 `json:"res_code,omitempty" xml:"res_code,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaDelayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResMsg = ""
	m.ResCode = 0
	m.IsSuccess = false
}

var poolTaobaoEticketMerchantMaDelayAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantMaDelayAPIResponse)
	},
}

// GetTaobaoEticketMerchantMaDelayAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantMaDelayAPIResponse
func GetTaobaoEticketMerchantMaDelayAPIResponse() *TaobaoEticketMerchantMaDelayAPIResponse {
	return poolTaobaoEticketMerchantMaDelayAPIResponse.Get().(*TaobaoEticketMerchantMaDelayAPIResponse)
}

// ReleaseTaobaoEticketMerchantMaDelayAPIResponse 将 TaobaoEticketMerchantMaDelayAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantMaDelayAPIResponse(v *TaobaoEticketMerchantMaDelayAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantMaDelayAPIResponse.Put(v)
}
