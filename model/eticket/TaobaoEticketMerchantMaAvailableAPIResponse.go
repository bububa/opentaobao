package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantMaAvailableAPIResponse 电子凭证核销前校验接口 API返回值
// taobao.eticket.merchant.ma.available
//
// 商家验码之前的调用接口，用来判断是否可以进行核销操作
type TaobaoEticketMerchantMaAvailableAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantMaAvailableAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaAvailableAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantMaAvailableAPIResponseModel).Reset()
}

// TaobaoEticketMerchantMaAvailableAPIResponseModel is 电子凭证核销前校验接口 成功返回结果
type TaobaoEticketMerchantMaAvailableAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_ma_available_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 子结果码
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// 子结果信息
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// 回复结果
	RespBody *ConsumeMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantMaAvailableAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantMaAvailableAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantMaAvailableAPIResponse)
	},
}

// GetTaobaoEticketMerchantMaAvailableAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantMaAvailableAPIResponse
func GetTaobaoEticketMerchantMaAvailableAPIResponse() *TaobaoEticketMerchantMaAvailableAPIResponse {
	return poolTaobaoEticketMerchantMaAvailableAPIResponse.Get().(*TaobaoEticketMerchantMaAvailableAPIResponse)
}

// ReleaseTaobaoEticketMerchantMaAvailableAPIResponse 将 TaobaoEticketMerchantMaAvailableAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantMaAvailableAPIResponse(v *TaobaoEticketMerchantMaAvailableAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantMaAvailableAPIResponse.Put(v)
}
