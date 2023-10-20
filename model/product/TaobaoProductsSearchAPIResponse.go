package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoProductsSearchAPIResponse 搜索产品信息 API返回值
// taobao.products.search
//
// 只有天猫商家发布商品时才需要用到，并非商品搜索api，当前暂不提供商品搜索api。&lt;br/&gt;二种方式搜索所有产品信息(二种至少传一种): &lt;br/&gt;
// 传入关键字q搜索&lt;br/&gt;
// 传入cid和props搜索&lt;br/&gt;
// 返回值支持:product_id,name,pic_path,cid,props,price,tsc&lt;br/&gt;
// 当用户指定了cid并且cid为垂直市场（3C电器城、鞋城）的类目id时，默认只返回小二确认&lt;br/&gt;的产品。如果用户没有指定cid，或cid为普通的类目，默认返回商家确认或小二确认的产&lt;br/&gt;品。如果用户自定了status字段，以指定的status类型为准。
// &lt;br/&gt;新增一：
//
//	传入suite_items_str 按规格搜索套装产品。
//	返回字段增加suite_items_str,is_suite_effecitve支持。
type TaobaoProductsSearchAPIResponse struct {
	model.CommonResponse
	TaobaoProductsSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoProductsSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoProductsSearchAPIResponseModel).Reset()
}

// TaobaoProductsSearchAPIResponseModel is 搜索产品信息 成功返回结果
type TaobaoProductsSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"products_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回具体信息为入参fields请求的字段信息
	Products []Product `json:"products,omitempty" xml:"products>product,omitempty"`
	// 结果总数
	TotalResults int64 `json:"total_results,omitempty" xml:"total_results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoProductsSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Products = m.Products[:0]
	m.TotalResults = 0
}

var poolTaobaoProductsSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoProductsSearchAPIResponse)
	},
}

// GetTaobaoProductsSearchAPIResponse 从 sync.Pool 获取 TaobaoProductsSearchAPIResponse
func GetTaobaoProductsSearchAPIResponse() *TaobaoProductsSearchAPIResponse {
	return poolTaobaoProductsSearchAPIResponse.Get().(*TaobaoProductsSearchAPIResponse)
}

// ReleaseTaobaoProductsSearchAPIResponse 将 TaobaoProductsSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoProductsSearchAPIResponse(v *TaobaoProductsSearchAPIResponse) {
	v.Reset()
	poolTaobaoProductsSearchAPIResponse.Put(v)
}
