package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductUpshelfAPIResponse 产品上架 API返回值
// tmall.supplychain.channel.product.upshelf
//
// 上架渠道产品
type TmallSupplychainChannelProductUpshelfAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductUpshelfAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductUpshelfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductUpshelfAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductUpshelfAPIResponseModel is 产品上架 成功返回结果
type TmallSupplychainChannelProductUpshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_upshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallSupplychainChannelProductUpshelfResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductUpshelfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductUpshelfAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductUpshelfAPIResponse)
	},
}

// GetTmallSupplychainChannelProductUpshelfAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductUpshelfAPIResponse
func GetTmallSupplychainChannelProductUpshelfAPIResponse() *TmallSupplychainChannelProductUpshelfAPIResponse {
	return poolTmallSupplychainChannelProductUpshelfAPIResponse.Get().(*TmallSupplychainChannelProductUpshelfAPIResponse)
}

// ReleaseTmallSupplychainChannelProductUpshelfAPIResponse 将 TmallSupplychainChannelProductUpshelfAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductUpshelfAPIResponse(v *TmallSupplychainChannelProductUpshelfAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductUpshelfAPIResponse.Put(v)
}
