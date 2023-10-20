package ju

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityIdentityStoreAPIResponse 用户信息存储 API返回值
// alibaba.jhs.community.identity.store
//
// 用户信息存储
type AlibabaJhsCommunityIdentityStoreAPIResponse struct {
	model.CommonResponse
	AlibabaJhsCommunityIdentityStoreAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityIdentityStoreAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJhsCommunityIdentityStoreAPIResponseModel).Reset()
}

// AlibabaJhsCommunityIdentityStoreAPIResponseModel is 用户信息存储 成功返回结果
type AlibabaJhsCommunityIdentityStoreAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_identity_store_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
	// 操作是否成功
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityIdentityStoreAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Error = ""
	m.Data = false
}

var poolAlibabaJhsCommunityIdentityStoreAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJhsCommunityIdentityStoreAPIResponse)
	},
}

// GetAlibabaJhsCommunityIdentityStoreAPIResponse 从 sync.Pool 获取 AlibabaJhsCommunityIdentityStoreAPIResponse
func GetAlibabaJhsCommunityIdentityStoreAPIResponse() *AlibabaJhsCommunityIdentityStoreAPIResponse {
	return poolAlibabaJhsCommunityIdentityStoreAPIResponse.Get().(*AlibabaJhsCommunityIdentityStoreAPIResponse)
}

// ReleaseAlibabaJhsCommunityIdentityStoreAPIResponse 将 AlibabaJhsCommunityIdentityStoreAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJhsCommunityIdentityStoreAPIResponse(v *AlibabaJhsCommunityIdentityStoreAPIResponse) {
	v.Reset()
	poolAlibabaJhsCommunityIdentityStoreAPIResponse.Put(v)
}
