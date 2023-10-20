package alilabs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabUserProfileGetAPIResponse 查询用户信息 API返回值
// alibaba.ailab.user.profile.get
//
// 提供天猫精灵用户头像、昵称的查询接口，供本田车载天猫精灵使用
type AlibabaAilabUserProfileGetAPIResponse struct {
	model.CommonResponse
	AlibabaAilabUserProfileGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAilabUserProfileGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAilabUserProfileGetAPIResponseModel).Reset()
}

// AlibabaAilabUserProfileGetAPIResponseModel is 查询用户信息 成功返回结果
type AlibabaAilabUserProfileGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailab_user_profile_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口model
	Result *AlibabaAilabUserProfileGetResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAilabUserProfileGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAilabUserProfileGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAilabUserProfileGetAPIResponse)
	},
}

// GetAlibabaAilabUserProfileGetAPIResponse 从 sync.Pool 获取 AlibabaAilabUserProfileGetAPIResponse
func GetAlibabaAilabUserProfileGetAPIResponse() *AlibabaAilabUserProfileGetAPIResponse {
	return poolAlibabaAilabUserProfileGetAPIResponse.Get().(*AlibabaAilabUserProfileGetAPIResponse)
}

// ReleaseAlibabaAilabUserProfileGetAPIResponse 将 AlibabaAilabUserProfileGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAilabUserProfileGetAPIResponse(v *AlibabaAilabUserProfileGetAPIResponse) {
	v.Reset()
	poolAlibabaAilabUserProfileGetAPIResponse.Put(v)
}
