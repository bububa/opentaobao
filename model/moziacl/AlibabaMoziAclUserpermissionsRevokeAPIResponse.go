package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclUserpermissionsRevokeAPIResponse 回收账户权限 API返回值
// alibaba.mozi.acl.userpermissions.revoke
//
// 调用此接口，会根据入参回收该账户下的该批权限点
type AlibabaMoziAclUserpermissionsRevokeAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclUserpermissionsRevokeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclUserpermissionsRevokeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclUserpermissionsRevokeAPIResponseModel).Reset()
}

// AlibabaMoziAclUserpermissionsRevokeAPIResponseModel is 回收账户权限 成功返回结果
type AlibabaMoziAclUserpermissionsRevokeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_userpermissions_revoke_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 回收账户的权限请求返回对象
	Result *RevokePermissionsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclUserpermissionsRevokeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziAclUserpermissionsRevokeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclUserpermissionsRevokeAPIResponse)
	},
}

// GetAlibabaMoziAclUserpermissionsRevokeAPIResponse 从 sync.Pool 获取 AlibabaMoziAclUserpermissionsRevokeAPIResponse
func GetAlibabaMoziAclUserpermissionsRevokeAPIResponse() *AlibabaMoziAclUserpermissionsRevokeAPIResponse {
	return poolAlibabaMoziAclUserpermissionsRevokeAPIResponse.Get().(*AlibabaMoziAclUserpermissionsRevokeAPIResponse)
}

// ReleaseAlibabaMoziAclUserpermissionsRevokeAPIResponse 将 AlibabaMoziAclUserpermissionsRevokeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclUserpermissionsRevokeAPIResponse(v *AlibabaMoziAclUserpermissionsRevokeAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclUserpermissionsRevokeAPIResponse.Put(v)
}
