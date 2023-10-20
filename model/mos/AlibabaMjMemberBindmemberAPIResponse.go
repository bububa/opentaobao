package mos

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjMemberBindmemberAPIResponse 绑定会员 API返回值
// alibaba.mj.member.bindmember
//
// 用于绑定喵街数字化会员
type AlibabaMjMemberBindmemberAPIResponse struct {
	model.CommonResponse
	AlibabaMjMemberBindmemberAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMjMemberBindmemberAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMjMemberBindmemberAPIResponseModel).Reset()
}

// AlibabaMjMemberBindmemberAPIResponseModel is 绑定会员 成功返回结果
type AlibabaMjMemberBindmemberAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_member_bindmember_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *SingleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMjMemberBindmemberAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMjMemberBindmemberAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMjMemberBindmemberAPIResponse)
	},
}

// GetAlibabaMjMemberBindmemberAPIResponse 从 sync.Pool 获取 AlibabaMjMemberBindmemberAPIResponse
func GetAlibabaMjMemberBindmemberAPIResponse() *AlibabaMjMemberBindmemberAPIResponse {
	return poolAlibabaMjMemberBindmemberAPIResponse.Get().(*AlibabaMjMemberBindmemberAPIResponse)
}

// ReleaseAlibabaMjMemberBindmemberAPIResponse 将 AlibabaMjMemberBindmemberAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMjMemberBindmemberAPIResponse(v *AlibabaMjMemberBindmemberAPIResponse) {
	v.Reset()
	poolAlibabaMjMemberBindmemberAPIResponse.Put(v)
}
