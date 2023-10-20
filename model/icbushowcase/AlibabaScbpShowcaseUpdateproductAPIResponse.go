package icbushowcase

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpShowcaseUpdateproductAPIResponse 替换橱窗商品 API返回值
// alibaba.scbp.showcase.updateproduct
//
// 替换橱窗商品
type AlibabaScbpShowcaseUpdateproductAPIResponse struct {
	model.CommonResponse
	AlibabaScbpShowcaseUpdateproductAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseUpdateproductAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpShowcaseUpdateproductAPIResponseModel).Reset()
}

// AlibabaScbpShowcaseUpdateproductAPIResponseModel is 替换橱窗商品 成功返回结果
type AlibabaScbpShowcaseUpdateproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_showcase_updateproduct_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否修改成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpShowcaseUpdateproductAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpShowcaseUpdateproductAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpShowcaseUpdateproductAPIResponse)
	},
}

// GetAlibabaScbpShowcaseUpdateproductAPIResponse 从 sync.Pool 获取 AlibabaScbpShowcaseUpdateproductAPIResponse
func GetAlibabaScbpShowcaseUpdateproductAPIResponse() *AlibabaScbpShowcaseUpdateproductAPIResponse {
	return poolAlibabaScbpShowcaseUpdateproductAPIResponse.Get().(*AlibabaScbpShowcaseUpdateproductAPIResponse)
}

// ReleaseAlibabaScbpShowcaseUpdateproductAPIResponse 将 AlibabaScbpShowcaseUpdateproductAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpShowcaseUpdateproductAPIResponse(v *AlibabaScbpShowcaseUpdateproductAPIResponse) {
	v.Reset()
	poolAlibabaScbpShowcaseUpdateproductAPIResponse.Put(v)
}
