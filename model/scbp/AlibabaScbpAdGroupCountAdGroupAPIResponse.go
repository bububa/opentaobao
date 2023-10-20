package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupCountAdGroupAPIResponse 统计adgroup数量 API返回值
// alibaba.scbp.ad.group.count.ad.group
//
// 统计adgroup数量
type AlibabaScbpAdGroupCountAdGroupAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupCountAdGroupAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupCountAdGroupAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupCountAdGroupAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupCountAdGroupAPIResponseModel is 统计adgroup数量 成功返回结果
type AlibabaScbpAdGroupCountAdGroupAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_count_ad_group_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupCountAdGroupAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = 0
}

var poolAlibabaScbpAdGroupCountAdGroupAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupCountAdGroupAPIResponse)
	},
}

// GetAlibabaScbpAdGroupCountAdGroupAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupCountAdGroupAPIResponse
func GetAlibabaScbpAdGroupCountAdGroupAPIResponse() *AlibabaScbpAdGroupCountAdGroupAPIResponse {
	return poolAlibabaScbpAdGroupCountAdGroupAPIResponse.Get().(*AlibabaScbpAdGroupCountAdGroupAPIResponse)
}

// ReleaseAlibabaScbpAdGroupCountAdGroupAPIResponse 将 AlibabaScbpAdGroupCountAdGroupAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupCountAdGroupAPIResponse(v *AlibabaScbpAdGroupCountAdGroupAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupCountAdGroupAPIResponse.Put(v)
}
