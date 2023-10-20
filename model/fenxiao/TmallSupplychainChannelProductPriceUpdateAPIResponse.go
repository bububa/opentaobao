package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductPriceUpdateAPIResponse 渠道价格更新接口 API返回值
// tmall.supplychain.channel.product.price.update
//
// 更新渠道产品价格
type TmallSupplychainChannelProductPriceUpdateAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductPriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductPriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductPriceUpdateAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductPriceUpdateAPIResponseModel is 渠道价格更新接口 成功返回结果
type TmallSupplychainChannelProductPriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_price_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *ResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductPriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductPriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductPriceUpdateAPIResponse)
	},
}

// GetTmallSupplychainChannelProductPriceUpdateAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductPriceUpdateAPIResponse
func GetTmallSupplychainChannelProductPriceUpdateAPIResponse() *TmallSupplychainChannelProductPriceUpdateAPIResponse {
	return poolTmallSupplychainChannelProductPriceUpdateAPIResponse.Get().(*TmallSupplychainChannelProductPriceUpdateAPIResponse)
}

// ReleaseTmallSupplychainChannelProductPriceUpdateAPIResponse 将 TmallSupplychainChannelProductPriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductPriceUpdateAPIResponse(v *TmallSupplychainChannelProductPriceUpdateAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductPriceUpdateAPIResponse.Put(v)
}
