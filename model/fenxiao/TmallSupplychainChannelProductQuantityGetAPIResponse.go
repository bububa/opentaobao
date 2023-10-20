package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductQuantityGetAPIResponse 渠道库存查询接口 API返回值
// tmall.supplychain.channel.product.quantity.get
//
// 渠道库存查询接口
type TmallSupplychainChannelProductQuantityGetAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductQuantityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductQuantityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductQuantityGetAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductQuantityGetAPIResponseModel is 渠道库存查询接口 成功返回结果
type TmallSupplychainChannelProductQuantityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_quantity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallSupplychainChannelProductQuantityGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductQuantityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductQuantityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductQuantityGetAPIResponse)
	},
}

// GetTmallSupplychainChannelProductQuantityGetAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductQuantityGetAPIResponse
func GetTmallSupplychainChannelProductQuantityGetAPIResponse() *TmallSupplychainChannelProductQuantityGetAPIResponse {
	return poolTmallSupplychainChannelProductQuantityGetAPIResponse.Get().(*TmallSupplychainChannelProductQuantityGetAPIResponse)
}

// ReleaseTmallSupplychainChannelProductQuantityGetAPIResponse 将 TmallSupplychainChannelProductQuantityGetAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductQuantityGetAPIResponse(v *TmallSupplychainChannelProductQuantityGetAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductQuantityGetAPIResponse.Put(v)
}
