package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecsGetAPIResponse 获取产品的规格信息 API返回值
// tmall.product.specs.get
//
// 按product_id或品牌下载产品规格，返回一组的产品规格信息。
type TmallProductSpecsGetAPIResponse struct {
	model.CommonResponse
	TmallProductSpecsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSpecsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSpecsGetAPIResponseModel).Reset()
}

// TmallProductSpecsGetAPIResponseModel is 获取产品的规格信息 成功返回结果
type TmallProductSpecsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_specs_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回一组产品规格信息。
	ProductSpecs []ProductSpec `json:"product_specs,omitempty" xml:"product_specs>product_spec,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSpecsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductSpecs = m.ProductSpecs[:0]
}

var poolTmallProductSpecsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSpecsGetAPIResponse)
	},
}

// GetTmallProductSpecsGetAPIResponse 从 sync.Pool 获取 TmallProductSpecsGetAPIResponse
func GetTmallProductSpecsGetAPIResponse() *TmallProductSpecsGetAPIResponse {
	return poolTmallProductSpecsGetAPIResponse.Get().(*TmallProductSpecsGetAPIResponse)
}

// ReleaseTmallProductSpecsGetAPIResponse 将 TmallProductSpecsGetAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSpecsGetAPIResponse(v *TmallProductSpecsGetAPIResponse) {
	v.Reset()
	poolTmallProductSpecsGetAPIResponse.Put(v)
}
