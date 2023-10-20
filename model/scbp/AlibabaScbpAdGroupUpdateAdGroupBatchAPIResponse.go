package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse 修改推广单元 API返回值
// alibaba.scbp.ad.group.update.ad.group.batch
//
// 修改推广单元
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponseModel is 修改推广单元 成功返回结果
type AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_update_ad_group_batch_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse)
	},
}

// GetAlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse
func GetAlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse() *AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse {
	return poolAlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse.Get().(*AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse)
}

// ReleaseAlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse 将 AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse(v *AlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupUpdateAdGroupBatchAPIResponse.Put(v)
}
