package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIdentityRescindfinishAPIResponse 取消确认 API返回值
// alibaba.member.identity.rescindfinish
//
// 取消确认
type AlibabaMemberIdentityRescindfinishAPIResponse struct {
	model.CommonResponse
	AlibabaMemberIdentityRescindfinishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberIdentityRescindfinishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberIdentityRescindfinishAPIResponseModel).Reset()
}

// AlibabaMemberIdentityRescindfinishAPIResponseModel is 取消确认 成功返回结果
type AlibabaMemberIdentityRescindfinishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_identity_rescindfinish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMemberIdentityRescindfinishTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberIdentityRescindfinishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMemberIdentityRescindfinishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberIdentityRescindfinishAPIResponse)
	},
}

// GetAlibabaMemberIdentityRescindfinishAPIResponse 从 sync.Pool 获取 AlibabaMemberIdentityRescindfinishAPIResponse
func GetAlibabaMemberIdentityRescindfinishAPIResponse() *AlibabaMemberIdentityRescindfinishAPIResponse {
	return poolAlibabaMemberIdentityRescindfinishAPIResponse.Get().(*AlibabaMemberIdentityRescindfinishAPIResponse)
}

// ReleaseAlibabaMemberIdentityRescindfinishAPIResponse 将 AlibabaMemberIdentityRescindfinishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberIdentityRescindfinishAPIResponse(v *AlibabaMemberIdentityRescindfinishAPIResponse) {
	v.Reset()
	poolAlibabaMemberIdentityRescindfinishAPIResponse.Put(v)
}
