package ju

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJhsCommunityWechatLoginAPIResponse 聚划算用增淘外社群登录 API返回值
// alibaba.jhs.community.wechat.login
//
// 聚划算用增淘外社群登录
type AlibabaJhsCommunityWechatLoginAPIResponse struct {
	model.CommonResponse
	AlibabaJhsCommunityWechatLoginAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityWechatLoginAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaJhsCommunityWechatLoginAPIResponseModel).Reset()
}

// AlibabaJhsCommunityWechatLoginAPIResponseModel is 聚划算用增淘外社群登录 成功返回结果
type AlibabaJhsCommunityWechatLoginAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_jhs_community_wechat_login_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 登录会话
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty" xml:"error,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaJhsCommunityWechatLoginAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.Error = ""
}

var poolAlibabaJhsCommunityWechatLoginAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaJhsCommunityWechatLoginAPIResponse)
	},
}

// GetAlibabaJhsCommunityWechatLoginAPIResponse 从 sync.Pool 获取 AlibabaJhsCommunityWechatLoginAPIResponse
func GetAlibabaJhsCommunityWechatLoginAPIResponse() *AlibabaJhsCommunityWechatLoginAPIResponse {
	return poolAlibabaJhsCommunityWechatLoginAPIResponse.Get().(*AlibabaJhsCommunityWechatLoginAPIResponse)
}

// ReleaseAlibabaJhsCommunityWechatLoginAPIResponse 将 AlibabaJhsCommunityWechatLoginAPIResponse 保存到 sync.Pool
func ReleaseAlibabaJhsCommunityWechatLoginAPIResponse(v *AlibabaJhsCommunityWechatLoginAPIResponse) {
	v.Reset()
	poolAlibabaJhsCommunityWechatLoginAPIResponse.Put(v)
}
