package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductPriceGetAPIResponse 渠道价格查询接口 API返回值
// tmall.supplychain.channel.product.price.get
//
// 渠道价格查询接口
type TmallSupplychainChannelProductPriceGetAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductPriceGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductPriceGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductPriceGetAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductPriceGetAPIResponseModel is 渠道价格查询接口 成功返回结果
type TmallSupplychainChannelProductPriceGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_price_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallSupplychainChannelProductPriceGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductPriceGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductPriceGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductPriceGetAPIResponse)
	},
}

// GetTmallSupplychainChannelProductPriceGetAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductPriceGetAPIResponse
func GetTmallSupplychainChannelProductPriceGetAPIResponse() *TmallSupplychainChannelProductPriceGetAPIResponse {
	return poolTmallSupplychainChannelProductPriceGetAPIResponse.Get().(*TmallSupplychainChannelProductPriceGetAPIResponse)
}

// ReleaseTmallSupplychainChannelProductPriceGetAPIResponse 将 TmallSupplychainChannelProductPriceGetAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductPriceGetAPIResponse(v *TmallSupplychainChannelProductPriceGetAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductPriceGetAPIResponse.Put(v)
}
