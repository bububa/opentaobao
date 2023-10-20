package eticket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoVmarketEticketCodesGetAPIResponse 电子凭证码列表查询 API返回值
// taobao.vmarket.eticket.codes.get
//
// 查询某个订单的所有码的列表
type TaobaoVmarketEticketCodesGetAPIResponse struct {
	model.CommonResponse
	TaobaoVmarketEticketCodesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketCodesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoVmarketEticketCodesGetAPIResponseModel).Reset()
}

// TaobaoVmarketEticketCodesGetAPIResponseModel is 电子凭证码列表查询 成功返回结果
type TaobaoVmarketEticketCodesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"vmarket_eticket_codes_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 电子凭证码列表
	EticketCodes []EticketCode `json:"eticket_codes,omitempty" xml:"eticket_codes>eticket_code,omitempty"`
	// 记录总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoVmarketEticketCodesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.EticketCodes = m.EticketCodes[:0]
	m.TotalResults = 0
}

var poolTaobaoVmarketEticketCodesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoVmarketEticketCodesGetAPIResponse)
	},
}

// GetTaobaoVmarketEticketCodesGetAPIResponse 从 sync.Pool 获取 TaobaoVmarketEticketCodesGetAPIResponse
func GetTaobaoVmarketEticketCodesGetAPIResponse() *TaobaoVmarketEticketCodesGetAPIResponse {
	return poolTaobaoVmarketEticketCodesGetAPIResponse.Get().(*TaobaoVmarketEticketCodesGetAPIResponse)
}

// ReleaseTaobaoVmarketEticketCodesGetAPIResponse 将 TaobaoVmarketEticketCodesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoVmarketEticketCodesGetAPIResponse(v *TaobaoVmarketEticketCodesGetAPIResponse) {
	v.Reset()
	poolTaobaoVmarketEticketCodesGetAPIResponse.Put(v)
}
