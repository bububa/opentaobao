package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoEticketMerchantTbmaGetAPIResponse 码商查询淘宝码接口 API返回值
// taobao.eticket.merchant.tbma.get
//
// 码商查询淘宝码接口
type TaobaoEticketMerchantTbmaGetAPIResponse struct {
	model.CommonResponse
	TaobaoEticketMerchantTbmaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantTbmaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoEticketMerchantTbmaGetAPIResponseModel).Reset()
}

// TaobaoEticketMerchantTbmaGetAPIResponseModel is 码商查询淘宝码接口 成功返回结果
type TaobaoEticketMerchantTbmaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"eticket_merchant_tbma_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// subCode
	RetCode string `json:"ret_code,omitempty" xml:"ret_code,omitempty"`
	// subMsg
	RetMsg string `json:"ret_msg,omitempty" xml:"ret_msg,omitempty"`
	// respBody
	RespBody *QueryTbMaCallbackResp `json:"resp_body,omitempty" xml:"resp_body,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoEticketMerchantTbmaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RetCode = ""
	m.RetMsg = ""
	m.RespBody = nil
}

var poolTaobaoEticketMerchantTbmaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoEticketMerchantTbmaGetAPIResponse)
	},
}

// GetTaobaoEticketMerchantTbmaGetAPIResponse 从 sync.Pool 获取 TaobaoEticketMerchantTbmaGetAPIResponse
func GetTaobaoEticketMerchantTbmaGetAPIResponse() *TaobaoEticketMerchantTbmaGetAPIResponse {
	return poolTaobaoEticketMerchantTbmaGetAPIResponse.Get().(*TaobaoEticketMerchantTbmaGetAPIResponse)
}

// ReleaseTaobaoEticketMerchantTbmaGetAPIResponse 将 TaobaoEticketMerchantTbmaGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoEticketMerchantTbmaGetAPIResponse(v *TaobaoEticketMerchantTbmaGetAPIResponse) {
	v.Reset()
	poolTaobaoEticketMerchantTbmaGetAPIResponse.Put(v)
}
