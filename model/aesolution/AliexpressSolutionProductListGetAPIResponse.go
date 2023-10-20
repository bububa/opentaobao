package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductListGetAPIResponse Get product list API返回值
// aliexpress.solution.product.list.get
//
// Get product list
type AliexpressSolutionProductListGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionProductListGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionProductListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionProductListGetAPIResponseModel).Reset()
}

// AliexpressSolutionProductListGetAPIResponseModel is Get product list 成功返回结果
type AliexpressSolutionProductListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ItemListResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionProductListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionProductListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionProductListGetAPIResponse)
	},
}

// GetAliexpressSolutionProductListGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionProductListGetAPIResponse
func GetAliexpressSolutionProductListGetAPIResponse() *AliexpressSolutionProductListGetAPIResponse {
	return poolAliexpressSolutionProductListGetAPIResponse.Get().(*AliexpressSolutionProductListGetAPIResponse)
}

// ReleaseAliexpressSolutionProductListGetAPIResponse 将 AliexpressSolutionProductListGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionProductListGetAPIResponse(v *AliexpressSolutionProductListGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionProductListGetAPIResponse.Put(v)
}
