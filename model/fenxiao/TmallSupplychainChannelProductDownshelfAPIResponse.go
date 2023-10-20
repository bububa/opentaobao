package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallSupplychainChannelProductDownshelfAPIResponse 产品下架 API返回值
// tmall.supplychain.channel.product.downshelf
//
// 产品下架
type TmallSupplychainChannelProductDownshelfAPIResponse struct {
	model.CommonResponse
	TmallSupplychainChannelProductDownshelfAPIResponseModel
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductDownshelfAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSupplychainChannelProductDownshelfAPIResponseModel).Reset()
}

// TmallSupplychainChannelProductDownshelfAPIResponseModel is 产品下架 成功返回结果
type TmallSupplychainChannelProductDownshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_downshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallSupplychainChannelProductDownshelfResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSupplychainChannelProductDownshelfAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSupplychainChannelProductDownshelfAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSupplychainChannelProductDownshelfAPIResponse)
	},
}

// GetTmallSupplychainChannelProductDownshelfAPIResponse 从 sync.Pool 获取 TmallSupplychainChannelProductDownshelfAPIResponse
func GetTmallSupplychainChannelProductDownshelfAPIResponse() *TmallSupplychainChannelProductDownshelfAPIResponse {
	return poolTmallSupplychainChannelProductDownshelfAPIResponse.Get().(*TmallSupplychainChannelProductDownshelfAPIResponse)
}

// ReleaseTmallSupplychainChannelProductDownshelfAPIResponse 将 TmallSupplychainChannelProductDownshelfAPIResponse 保存到 sync.Pool
func ReleaseTmallSupplychainChannelProductDownshelfAPIResponse(v *TmallSupplychainChannelProductDownshelfAPIResponse) {
	v.Reset()
	poolTmallSupplychainChannelProductDownshelfAPIResponse.Put(v)
}
