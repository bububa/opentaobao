package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleAddAPIResponse 新增一个角色 API返回值
// alibaba.mozi.acl.role.add
//
// 新增一个角色
type AlibabaMoziAclRoleAddAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclRoleAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclRoleAddAPIResponseModel).Reset()
}

// AlibabaMoziAclRoleAddAPIResponseModel is 新增一个角色 成功返回结果
type AlibabaMoziAclRoleAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_role_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 创建角色结果对象
	Result *CreateRoleResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziAclRoleAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclRoleAddAPIResponse)
	},
}

// GetAlibabaMoziAclRoleAddAPIResponse 从 sync.Pool 获取 AlibabaMoziAclRoleAddAPIResponse
func GetAlibabaMoziAclRoleAddAPIResponse() *AlibabaMoziAclRoleAddAPIResponse {
	return poolAlibabaMoziAclRoleAddAPIResponse.Get().(*AlibabaMoziAclRoleAddAPIResponse)
}

// ReleaseAlibabaMoziAclRoleAddAPIResponse 将 AlibabaMoziAclRoleAddAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclRoleAddAPIResponse(v *AlibabaMoziAclRoleAddAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclRoleAddAPIResponse.Put(v)
}
