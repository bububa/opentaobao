package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdGroupFindAdGroupAPIResponse 查询推广组 API返回值
// alibaba.scbp.ad.group.find.ad.group
//
// 查询推广组
type AlibabaScbpAdGroupFindAdGroupAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdGroupFindAdGroupAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupFindAdGroupAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdGroupFindAdGroupAPIResponseModel).Reset()
}

// AlibabaScbpAdGroupFindAdGroupAPIResponseModel is 查询推广组 成功返回结果
type AlibabaScbpAdGroupFindAdGroupAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_find_ad_group_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *PageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdGroupFindAdGroupAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaScbpAdGroupFindAdGroupAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdGroupFindAdGroupAPIResponse)
	},
}

// GetAlibabaScbpAdGroupFindAdGroupAPIResponse 从 sync.Pool 获取 AlibabaScbpAdGroupFindAdGroupAPIResponse
func GetAlibabaScbpAdGroupFindAdGroupAPIResponse() *AlibabaScbpAdGroupFindAdGroupAPIResponse {
	return poolAlibabaScbpAdGroupFindAdGroupAPIResponse.Get().(*AlibabaScbpAdGroupFindAdGroupAPIResponse)
}

// ReleaseAlibabaScbpAdGroupFindAdGroupAPIResponse 将 AlibabaScbpAdGroupFindAdGroupAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdGroupFindAdGroupAPIResponse(v *AlibabaScbpAdGroupFindAdGroupAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdGroupFindAdGroupAPIResponse.Put(v)
}
