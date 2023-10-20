package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecAddAPIResponse 添加产品规格 API返回值
// tmall.product.spec.add
//
// 增加产品规格
type TmallProductSpecAddAPIResponse struct {
	model.CommonResponse
	TmallProductSpecAddAPIResponseModel
}

// Reset 清空结构体
func (m *TmallProductSpecAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallProductSpecAddAPIResponseModel).Reset()
}

// TmallProductSpecAddAPIResponseModel is 添加产品规格 成功返回结果
type TmallProductSpecAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_spec_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 产品规格对象
	ProductSpec *ProductSpec `json:"product_spec,omitempty" xml:"product_spec,omitempty"`
}

// Reset 清空结构体
func (m *TmallProductSpecAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ProductSpec = nil
}

var poolTmallProductSpecAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallProductSpecAddAPIResponse)
	},
}

// GetTmallProductSpecAddAPIResponse 从 sync.Pool 获取 TmallProductSpecAddAPIResponse
func GetTmallProductSpecAddAPIResponse() *TmallProductSpecAddAPIResponse {
	return poolTmallProductSpecAddAPIResponse.Get().(*TmallProductSpecAddAPIResponse)
}

// ReleaseTmallProductSpecAddAPIResponse 将 TmallProductSpecAddAPIResponse 保存到 sync.Pool
func ReleaseTmallProductSpecAddAPIResponse(v *TmallProductSpecAddAPIResponse) {
	v.Reset()
	poolTmallProductSpecAddAPIResponse.Put(v)
}
