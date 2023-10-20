package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductsGetAPIResponse 获取产品列表 API返回值
// taobao.products.get
//
// 根据淘宝会员帐号搜索所有产品信息，推荐使用taobao.products.search
// 注意：支持分页，每页最多返回100条,默认值为40,页码从1开始，默认为第一页
type TaobaoProductsGetAPIResponse struct {
	model.CommonResponse
	TaobaoProductsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoProductsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoProductsGetAPIResponseModel).Reset()
}

// TaobaoProductsGetAPIResponseModel is 获取产品列表 成功返回结果
type TaobaoProductsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"products_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回具体信息为入参fields请求的字段信息
	Products []Product `json:"products,omitempty" xml:"products>product,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoProductsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Products = m.Products[:0]
}

var poolTaobaoProductsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoProductsGetAPIResponse)
	},
}

// GetTaobaoProductsGetAPIResponse 从 sync.Pool 获取 TaobaoProductsGetAPIResponse
func GetTaobaoProductsGetAPIResponse() *TaobaoProductsGetAPIResponse {
	return poolTaobaoProductsGetAPIResponse.Get().(*TaobaoProductsGetAPIResponse)
}

// ReleaseTaobaoProductsGetAPIResponse 将 TaobaoProductsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoProductsGetAPIResponse(v *TaobaoProductsGetAPIResponse) {
	v.Reset()
	poolTaobaoProductsGetAPIResponse.Put(v)
}
