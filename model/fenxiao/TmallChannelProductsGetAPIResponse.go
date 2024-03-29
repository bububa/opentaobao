package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallChannelProductsGetAPIResponse 查询供应商的产品数据 API返回值
// tmall.channel.products.get
//
// 查询供应商的产品数据。
//
// * 入参传入pids将优先查询，即只按这个条件查询。
// *入参传入sku_number将优先查询(没有传入pids)，即只按这个条件查询(最多显示50条)
// * 入参fields传skus将查询sku的数据，不传该参数默认不查询，返回产品的其它信息。
// * 入参fields传入images将查询多图数据，不传只返回主图数据。
// * 入参fields仅对传入pids生效（只有按ID查询时，才能查询额外的数据）
// * 查询结果按照产品发布时间倒序，即时间近的数据在前。
// * 传入channel 渠道，会只返回相应渠道的产品
type TmallChannelProductsGetAPIResponse struct {
	model.CommonResponse
	TmallChannelProductsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallChannelProductsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallChannelProductsGetAPIResponseModel).Reset()
}

// TmallChannelProductsGetAPIResponseModel is 查询供应商的产品数据 成功返回结果
type TmallChannelProductsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_channel_products_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品对象记录集
	Products []TopProductDo `json:"products,omitempty" xml:"products>top_product_do,omitempty"`
	// 查询结果记录数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TmallChannelProductsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Products = m.Products[:0]
	m.TotalResults = 0
}

var poolTmallChannelProductsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallChannelProductsGetAPIResponse)
	},
}

// GetTmallChannelProductsGetAPIResponse 从 sync.Pool 获取 TmallChannelProductsGetAPIResponse
func GetTmallChannelProductsGetAPIResponse() *TmallChannelProductsGetAPIResponse {
	return poolTmallChannelProductsGetAPIResponse.Get().(*TmallChannelProductsGetAPIResponse)
}

// ReleaseTmallChannelProductsGetAPIResponse 将 TmallChannelProductsGetAPIResponse 保存到 sync.Pool
func ReleaseTmallChannelProductsGetAPIResponse(v *TmallChannelProductsGetAPIResponse) {
	v.Reset()
	poolTmallChannelProductsGetAPIResponse.Put(v)
}
