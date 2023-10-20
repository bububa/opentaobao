package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeConsigngoodsAPIResponse 卖家发货 API返回值
// tmall.exchange.consigngoods
//
// 卖家发货
type TmallExchangeConsigngoodsAPIResponse struct {
	model.CommonResponse
	TmallExchangeConsigngoodsAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeConsigngoodsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeConsigngoodsAPIResponseModel).Reset()
}

// TmallExchangeConsigngoodsAPIResponseModel is 卖家发货 成功返回结果
type TmallExchangeConsigngoodsAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_consigngoods_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RefundBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeConsigngoodsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeConsigngoodsAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeConsigngoodsAPIResponse)
	},
}

// GetTmallExchangeConsigngoodsAPIResponse 从 sync.Pool 获取 TmallExchangeConsigngoodsAPIResponse
func GetTmallExchangeConsigngoodsAPIResponse() *TmallExchangeConsigngoodsAPIResponse {
	return poolTmallExchangeConsigngoodsAPIResponse.Get().(*TmallExchangeConsigngoodsAPIResponse)
}

// ReleaseTmallExchangeConsigngoodsAPIResponse 将 TmallExchangeConsigngoodsAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeConsigngoodsAPIResponse(v *TmallExchangeConsigngoodsAPIResponse) {
	v.Reset()
	poolTmallExchangeConsigngoodsAPIResponse.Put(v)
}
