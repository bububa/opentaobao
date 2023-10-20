package alimember

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIdentitySignfinishAPIResponse 签约确认 API返回值
// alibaba.member.identity.signfinish
//
// 签约确认
type AlibabaMemberIdentitySignfinishAPIResponse struct {
	model.CommonResponse
	AlibabaMemberIdentitySignfinishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMemberIdentitySignfinishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMemberIdentitySignfinishAPIResponseModel).Reset()
}

// AlibabaMemberIdentitySignfinishAPIResponseModel is 签约确认 成功返回结果
type AlibabaMemberIdentitySignfinishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_member_identity_signfinish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaMemberIdentitySignfinishTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMemberIdentitySignfinishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMemberIdentitySignfinishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMemberIdentitySignfinishAPIResponse)
	},
}

// GetAlibabaMemberIdentitySignfinishAPIResponse 从 sync.Pool 获取 AlibabaMemberIdentitySignfinishAPIResponse
func GetAlibabaMemberIdentitySignfinishAPIResponse() *AlibabaMemberIdentitySignfinishAPIResponse {
	return poolAlibabaMemberIdentitySignfinishAPIResponse.Get().(*AlibabaMemberIdentitySignfinishAPIResponse)
}

// ReleaseAlibabaMemberIdentitySignfinishAPIResponse 将 AlibabaMemberIdentitySignfinishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMemberIdentitySignfinishAPIResponse(v *AlibabaMemberIdentitySignfinishAPIResponse) {
	v.Reset()
	poolAlibabaMemberIdentitySignfinishAPIResponse.Put(v)
}
