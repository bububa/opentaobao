package exchange

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallExchangeReturngoodsRefuseAPIResponse 卖家拒绝确认收货 API返回值
// tmall.exchange.returngoods.refuse
//
// 卖家拒绝买家换货申请
type TmallExchangeReturngoodsRefuseAPIResponse struct {
	model.CommonResponse
	TmallExchangeReturngoodsRefuseAPIResponseModel
}

// Reset 清空结构体
func (m *TmallExchangeReturngoodsRefuseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallExchangeReturngoodsRefuseAPIResponseModel).Reset()
}

// TmallExchangeReturngoodsRefuseAPIResponseModel is 卖家拒绝确认收货 成功返回结果
type TmallExchangeReturngoodsRefuseAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_exchange_returngoods_refuse_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ExchangeBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallExchangeReturngoodsRefuseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallExchangeReturngoodsRefuseAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallExchangeReturngoodsRefuseAPIResponse)
	},
}

// GetTmallExchangeReturngoodsRefuseAPIResponse 从 sync.Pool 获取 TmallExchangeReturngoodsRefuseAPIResponse
func GetTmallExchangeReturngoodsRefuseAPIResponse() *TmallExchangeReturngoodsRefuseAPIResponse {
	return poolTmallExchangeReturngoodsRefuseAPIResponse.Get().(*TmallExchangeReturngoodsRefuseAPIResponse)
}

// ReleaseTmallExchangeReturngoodsRefuseAPIResponse 将 TmallExchangeReturngoodsRefuseAPIResponse 保存到 sync.Pool
func ReleaseTmallExchangeReturngoodsRefuseAPIResponse(v *TmallExchangeReturngoodsRefuseAPIResponse) {
	v.Reset()
	poolTmallExchangeReturngoodsRefuseAPIResponse.Put(v)
}
