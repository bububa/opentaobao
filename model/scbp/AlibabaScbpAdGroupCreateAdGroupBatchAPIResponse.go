package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse 创建推广单元 API返回值
// alibaba.scbp.ad.group.create.ad.group.batch
//
// 创建推广单元
type AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupCreateAdGroupBatchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupCreateAdGroupBatchAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupCreateAdGroupBatchAPIResponseModel is 创建推广单元 成功返回结果
type AlibabaScbpAdGroupCreateAdGroupBatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_create_ad_group_batch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupCreateAdGroupBatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaScbpAdGroupCreateAdGroupBatchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse)
	},
}

// GetAlibabaScbpAdGroupCreateAdGroupBatchAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse
func GetAlibabaScbpAdGroupCreateAdGroupBatchAPIResponse() *AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse {
	return poolAlibabaScbpAdGroupCreateAdGroupBatchAPIResponse.Get().(*AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse)
}

// ReleaseAlibabaScbpAdGroupCreateAdGroupBatchAPIResponse 将 AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupCreateAdGroupBatchAPIResponse(v *AlibabaScbpAdGroupCreateAdGroupBatchAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupCreateAdGroupBatchAPIResponse.Put(v)
}
