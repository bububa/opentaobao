package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductPreferentialUpdateAPIResponse 设置P4P产品优先推广状态 API返回值
// alibaba.scbp.product.preferential.update
//
// 设置P4P产品优先推广状态
type AlibabaScbpProductPreferentialUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpProductPreferentialUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpProductPreferentialUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpProductPreferentialUpdateAPIResponseModel).Reset()
}

// AlibabaScbpProductPreferentialUpdateAPIResponseModel is 设置P4P产品优先推广状态 成功返回结果
type AlibabaScbpProductPreferentialUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_product_preferential_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设置成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpProductPreferentialUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaScbpProductPreferentialUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpProductPreferentialUpdateAPIResponse)
	},
}

// GetAlibabaScbpProductPreferentialUpdateAPIResponse 从 sync.Pool 获取 AlibabaScbpProductPreferentialUpdateAPIResponse
func GetAlibabaScbpProductPreferentialUpdateAPIResponse() *AlibabaScbpProductPreferentialUpdateAPIResponse {
	return poolAlibabaScbpProductPreferentialUpdateAPIResponse.Get().(*AlibabaScbpProductPreferentialUpdateAPIResponse)
}

// ReleaseAlibabaScbpProductPreferentialUpdateAPIResponse 将 AlibabaScbpProductPreferentialUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpProductPreferentialUpdateAPIResponse(v *AlibabaScbpProductPreferentialUpdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpProductPreferentialUpdateAPIResponse.Put(v)
}
