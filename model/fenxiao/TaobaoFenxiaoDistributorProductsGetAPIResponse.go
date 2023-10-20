package fenxiao

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFenxiaoDistributorProductsGetAPIResponse 分销商查询产品信息 API返回值
// taobao.fenxiao.distributor.products.get
//
// 分销商查询供应商产品信息
type TaobaoFenxiaoDistributorProductsGetAPIResponse struct {
	model.CommonResponse
	TaobaoFenxiaoDistributorProductsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDistributorProductsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFenxiaoDistributorProductsGetAPIResponseModel).Reset()
}

// TaobaoFenxiaoDistributorProductsGetAPIResponseModel is 分销商查询产品信息 成功返回结果
type TaobaoFenxiaoDistributorProductsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"fenxiao_distributor_products_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品对象记录集。返回 FenxiaoProduct 包含的字段信息。
	Products []FenxiaoProduct `json:"products,omitempty" xml:"products>fenxiao_product,omitempty"`
	// 是否存在下一页
	HasNext bool `json:"has_next,omitempty" xml:"has_next,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFenxiaoDistributorProductsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Products = m.Products[:0]
	m.HasNext = false
}

var poolTaobaoFenxiaoDistributorProductsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFenxiaoDistributorProductsGetAPIResponse)
	},
}

// GetTaobaoFenxiaoDistributorProductsGetAPIResponse 从 sync.Pool 获取 TaobaoFenxiaoDistributorProductsGetAPIResponse
func GetTaobaoFenxiaoDistributorProductsGetAPIResponse() *TaobaoFenxiaoDistributorProductsGetAPIResponse {
	return poolTaobaoFenxiaoDistributorProductsGetAPIResponse.Get().(*TaobaoFenxiaoDistributorProductsGetAPIResponse)
}

// ReleaseTaobaoFenxiaoDistributorProductsGetAPIResponse 将 TaobaoFenxiaoDistributorProductsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFenxiaoDistributorProductsGetAPIResponse(v *TaobaoFenxiaoDistributorProductsGetAPIResponse) {
	v.Reset()
	poolTaobaoFenxiaoDistributorProductsGetAPIResponse.Put(v)
}
