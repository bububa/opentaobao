package tmallchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelProductsQueryAPIResponse 渠道中心-查询产品列表 API返回值
// tmall.channel.products.query
//
// 渠道中心，供应商查询其产品数据，返回同时符合所有查询条件的产品信息
type TmallChannelProductsQueryAPIResponse struct {
	model.CommonResponse
	TmallChannelProductsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TmallChannelProductsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallChannelProductsQueryAPIResponseModel).Reset()
}

// TmallChannelProductsQueryAPIResponseModel is 渠道中心-查询产品列表 成功返回结果
type TmallChannelProductsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_products_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallChannelProductsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallChannelProductsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallChannelProductsQueryAPIResponse)
	},
}

// GetTmallChannelProductsQueryAPIResponse 从 sync.Pool 获取 TmallChannelProductsQueryAPIResponse
func GetTmallChannelProductsQueryAPIResponse() *TmallChannelProductsQueryAPIResponse {
	return poolTmallChannelProductsQueryAPIResponse.Get().(*TmallChannelProductsQueryAPIResponse)
}

// ReleaseTmallChannelProductsQueryAPIResponse 将 TmallChannelProductsQueryAPIResponse 保存到 sync.Pool
func ReleaseTmallChannelProductsQueryAPIResponse(v *TmallChannelProductsQueryAPIResponse) {
	v.Reset()
	poolTmallChannelProductsQueryAPIResponse.Put(v)
}
