package icbushowcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseDeleteproductAPIResponse 批量删除橱窗商品 API返回值
// alibaba.scbp.showcase.deleteproduct
//
// 批量删除橱窗商品
type AlibabaScbpShowcaseDeleteproductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpShowcaseDeleteproductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseDeleteproductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpShowcaseDeleteproductAPIResponseModel).Reset()
}

// AlibabaScbpShowcaseDeleteproductAPIResponseModel is 批量删除橱窗商品 成功返回结果
type AlibabaScbpShowcaseDeleteproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_deleteproduct_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseDeleteproductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpShowcaseDeleteproductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpShowcaseDeleteproductAPIResponse)
	},
}

// GetAlibabaScbpShowcaseDeleteproductAPIResponse 从 sync.Pool 获取 AlibabaScbpShowcaseDeleteproductAPIResponse
func GetAlibabaScbpShowcaseDeleteproductAPIResponse() *AlibabaScbpShowcaseDeleteproductAPIResponse {
	return poolAlibabaScbpShowcaseDeleteproductAPIResponse.Get().(*AlibabaScbpShowcaseDeleteproductAPIResponse)
}

// ReleaseAlibabaScbpShowcaseDeleteproductAPIResponse 将 AlibabaScbpShowcaseDeleteproductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpShowcaseDeleteproductAPIResponse(v *AlibabaScbpShowcaseDeleteproductAPIResponse) {
	v.Reset()
	poolAlibabaScbpShowcaseDeleteproductAPIResponse.Put(v)
}
