package icbushowcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseAddproductAPIResponse 批量添加橱窗商品 API返回值
// alibaba.scbp.showcase.addproduct
//
// 批量添加商品到橱窗
type AlibabaScbpShowcaseAddproductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpShowcaseAddproductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseAddproductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpShowcaseAddproductAPIResponseModel).Reset()
}

// AlibabaScbpShowcaseAddproductAPIResponseModel is 批量添加橱窗商品 成功返回结果
type AlibabaScbpShowcaseAddproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_addproduct_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否添加成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseAddproductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpShowcaseAddproductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpShowcaseAddproductAPIResponse)
	},
}

// GetAlibabaScbpShowcaseAddproductAPIResponse 从 sync.Pool 获取 AlibabaScbpShowcaseAddproductAPIResponse
func GetAlibabaScbpShowcaseAddproductAPIResponse() *AlibabaScbpShowcaseAddproductAPIResponse {
	return poolAlibabaScbpShowcaseAddproductAPIResponse.Get().(*AlibabaScbpShowcaseAddproductAPIResponse)
}

// ReleaseAlibabaScbpShowcaseAddproductAPIResponse 将 AlibabaScbpShowcaseAddproductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpShowcaseAddproductAPIResponse(v *AlibabaScbpShowcaseAddproductAPIResponse) {
	v.Reset()
	poolAlibabaScbpShowcaseAddproductAPIResponse.Put(v)
}
