package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeReturngoodsAgreeAPIResponse 卖家确认收货 API返回值
// tmall.exchange.returngoods.agree
//
// 卖家确认收货
type TmallExchangeReturngoodsAgreeAPIResponse struct {
	model.CommonResponse
	TmallExchangeReturngoodsAgreeAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeReturngoodsAgreeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeReturngoodsAgreeAPIResponseModel).Reset()
}

// TmallExchangeReturngoodsAgreeAPIResponseModel is 卖家确认收货 成功返回结果
type TmallExchangeReturngoodsAgreeAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_returngoods_agree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeReturngoodsAgreeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeReturngoodsAgreeAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeReturngoodsAgreeAPIResponse)
	},
}

// GetTmallExchangeReturngoodsAgreeAPIResponse 从 sync.Pool 获取 TmallExchangeReturngoodsAgreeAPIResponse
func GetTmallExchangeReturngoodsAgreeAPIResponse() *TmallExchangeReturngoodsAgreeAPIResponse {
	return poolTmallExchangeReturngoodsAgreeAPIResponse.Get().(*TmallExchangeReturngoodsAgreeAPIResponse)
}

// ReleaseTmallExchangeReturngoodsAgreeAPIResponse 将 TmallExchangeReturngoodsAgreeAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeReturngoodsAgreeAPIResponse(v *TmallExchangeReturngoodsAgreeAPIResponse) {
	v.Reset()
	poolTmallExchangeReturngoodsAgreeAPIResponse.Put(v)
}
