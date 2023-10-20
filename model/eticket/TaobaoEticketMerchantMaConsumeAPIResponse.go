package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaConsumeAPIResponse 电子凭证核销接口 API返回值
// taobao.eticket.merchant.ma.consume
//
// 电子凭证核销接口
type TaobaoEticketMerchantMaConsumeAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaConsumeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaConsumeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantMaConsumeAPIResponseModel).Reset()
}

// TaobaoEticketMerchantMaConsumeAPIResponseModel is 电子凭证核销接口 成功返回结果
type TaobaoEticketMerchantMaConsumeAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_consume_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 系统自动生成
	RespBody *ConsumeMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaConsumeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantMaConsumeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantMaConsumeAPIResponse)
	},
}

// GetTaobaoEticketMerchantMaConsumeAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantMaConsumeAPIResponse
func GetTaobaoEticketMerchantMaConsumeAPIResponse() *TaobaoEticketMerchantMaConsumeAPIResponse {
	return poolTaobaoEticketMerchantMaConsumeAPIResponse.Get().(*TaobaoEticketMerchantMaConsumeAPIResponse)
}

// ReleaseTaobaoEticketMerchantMaConsumeAPIResponse 将 TaobaoEticketMerchantMaConsumeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantMaConsumeAPIResponse(v *TaobaoEticketMerchantMaConsumeAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantMaConsumeAPIResponse.Put(v)
}
