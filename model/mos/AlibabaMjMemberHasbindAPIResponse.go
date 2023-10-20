package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMemberHasbindAPIResponse 喵街会员是否绑定 API返回值
// alibaba.mj.member.hasbind
//
// 喵街检测用户是否为数字化会员
type AlibabaMjMemberHasbindAPIResponse struct {
	model.CommonResponse
	AlibabaMjMemberHasbindAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjMemberHasbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjMemberHasbindAPIResponseModel).Reset()
}

// AlibabaMjMemberHasbindAPIResponseModel is 喵街会员是否绑定 成功返回结果
type AlibabaMjMemberHasbindAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_member_hasbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjMemberHasbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMjMemberHasbindAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjMemberHasbindAPIResponse)
	},
}

// GetAlibabaMjMemberHasbindAPIResponse 从 sync.Pool 获取 AlibabaMjMemberHasbindAPIResponse
func GetAlibabaMjMemberHasbindAPIResponse() *AlibabaMjMemberHasbindAPIResponse {
	return poolAlibabaMjMemberHasbindAPIResponse.Get().(*AlibabaMjMemberHasbindAPIResponse)
}

// ReleaseAlibabaMjMemberHasbindAPIResponse 将 AlibabaMjMemberHasbindAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjMemberHasbindAPIResponse(v *AlibabaMjMemberHasbindAPIResponse) {
	v.Reset()
	poolAlibabaMjMemberHasbindAPIResponse.Put(v)
}
