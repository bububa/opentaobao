package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclUserrolesRevokeAPIResponse 回收账户被授予的角色接口 API返回值
// alibaba.mozi.acl.userroles.revoke
//
// 调用此接口，会根据入参回收该账户下的该批角色
type AlibabaMoziAclUserrolesRevokeAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclUserrolesRevokeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclUserrolesRevokeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclUserrolesRevokeAPIResponseModel).Reset()
}

// AlibabaMoziAclUserrolesRevokeAPIResponseModel is 回收账户被授予的角色接口 成功返回结果
type AlibabaMoziAclUserrolesRevokeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_userroles_revoke_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回收角色结果对象
	Result *RevokeRolesResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclUserrolesRevokeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziAclUserrolesRevokeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclUserrolesRevokeAPIResponse)
	},
}

// GetAlibabaMoziAclUserrolesRevokeAPIResponse 从 sync.Pool 获取 AlibabaMoziAclUserrolesRevokeAPIResponse
func GetAlibabaMoziAclUserrolesRevokeAPIResponse() *AlibabaMoziAclUserrolesRevokeAPIResponse {
	return poolAlibabaMoziAclUserrolesRevokeAPIResponse.Get().(*AlibabaMoziAclUserrolesRevokeAPIResponse)
}

// ReleaseAlibabaMoziAclUserrolesRevokeAPIResponse 将 AlibabaMoziAclUserrolesRevokeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclUserrolesRevokeAPIResponse(v *AlibabaMoziAclUserrolesRevokeAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclUserrolesRevokeAPIResponse.Put(v)
}
