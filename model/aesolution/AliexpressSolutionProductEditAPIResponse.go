package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductEditAPIResponse Edit Product API API返回值
// aliexpress.solution.product.edit
//
// API for editing product, customized for Oversea merchants. Most of the input fields of this API is similar with aliexpress.solution.product.post. For editing, just fill in the fields you would like to edit, leave other fields to be blank.
type AliexpressSolutionProductEditAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionProductEditAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionProductEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionProductEditAPIResponseModel).Reset()
}

// AliexpressSolutionProductEditAPIResponseModel is Edit Product API 成功返回结果
type AliexpressSolutionProductEditAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_product_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PostItemResponseDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionProductEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAliexpressSolutionProductEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionProductEditAPIResponse)
	},
}

// GetAliexpressSolutionProductEditAPIResponse 从 sync.Pool 获取 AliexpressSolutionProductEditAPIResponse
func GetAliexpressSolutionProductEditAPIResponse() *AliexpressSolutionProductEditAPIResponse {
	return poolAliexpressSolutionProductEditAPIResponse.Get().(*AliexpressSolutionProductEditAPIResponse)
}

// ReleaseAliexpressSolutionProductEditAPIResponse 将 AliexpressSolutionProductEditAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionProductEditAPIResponse(v *AliexpressSolutionProductEditAPIResponse) {
	v.Reset()
	poolAliexpressSolutionProductEditAPIResponse.Put(v)
}
