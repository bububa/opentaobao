package alihealthmdeer

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthMdeerScienceDeletearticleAPIResponse 文章删除 API返回值
// alibaba.alihealth.mdeer.science.deletearticle
//
// 三方同步文章删除
type AlibabaAlihealthMdeerScienceDeletearticleAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthMdeerScienceDeletearticleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel).Reset()
}

// AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel is 文章删除 成功返回结果
type AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_mdeer_science_deletearticle_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 错误code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 是否删除成功
	Model bool `json:"model,omitempty" xml:"model,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthMdeerScienceDeletearticleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MsgInfo = ""
	m.MsgCode = ""
	m.Model = false
}

var poolAlibabaAlihealthMdeerScienceDeletearticleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthMdeerScienceDeletearticleAPIResponse)
	},
}

// GetAlibabaAlihealthMdeerScienceDeletearticleAPIResponse 从 sync.Pool 获取 AlibabaAlihealthMdeerScienceDeletearticleAPIResponse
func GetAlibabaAlihealthMdeerScienceDeletearticleAPIResponse() *AlibabaAlihealthMdeerScienceDeletearticleAPIResponse {
	return poolAlibabaAlihealthMdeerScienceDeletearticleAPIResponse.Get().(*AlibabaAlihealthMdeerScienceDeletearticleAPIResponse)
}

// ReleaseAlibabaAlihealthMdeerScienceDeletearticleAPIResponse 将 AlibabaAlihealthMdeerScienceDeletearticleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthMdeerScienceDeletearticleAPIResponse(v *AlibabaAlihealthMdeerScienceDeletearticleAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthMdeerScienceDeletearticleAPIResponse.Put(v)
}
