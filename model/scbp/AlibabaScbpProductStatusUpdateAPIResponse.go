package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpProductStatusUpdateAPIResponse 修改P4P产品推广状态 API返回值
// alibaba.scbp.product.status.update
//
// 修改P4P产品推广状态
type AlibabaScbpProductStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpProductStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpProductStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpProductStatusUpdateAPIResponseModel).Reset()
}

// AlibabaScbpProductStatusUpdateAPIResponseModel is 修改P4P产品推广状态 成功返回结果
type AlibabaScbpProductStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_product_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 实际修改的产品数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpProductStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpProductStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpProductStatusUpdateAPIResponse)
	},
}

// GetAlibabaScbpProductStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaScbpProductStatusUpdateAPIResponse
func GetAlibabaScbpProductStatusUpdateAPIResponse() *AlibabaScbpProductStatusUpdateAPIResponse {
	return poolAlibabaScbpProductStatusUpdateAPIResponse.Get().(*AlibabaScbpProductStatusUpdateAPIResponse)
}

// ReleaseAlibabaScbpProductStatusUpdateAPIResponse 将 AlibabaScbpProductStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpProductStatusUpdateAPIResponse(v *AlibabaScbpProductStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaScbpProductStatusUpdateAPIResponse.Put(v)
}
