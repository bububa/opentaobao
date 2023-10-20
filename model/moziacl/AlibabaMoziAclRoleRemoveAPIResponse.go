package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclRoleRemoveAPIResponse 删除角色 API返回值
// alibaba.mozi.acl.role.remove
//
// 根据传入的角色code、租户id，删除租户内对应的角色
type AlibabaMoziAclRoleRemoveAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclRoleRemoveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleRemoveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclRoleRemoveAPIResponseModel).Reset()
}

// AlibabaMoziAclRoleRemoveAPIResponseModel is 删除角色 成功返回结果
type AlibabaMoziAclRoleRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_role_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 无值
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 请求id
	MoziRequestId string `json:"mozi_request_id,omitempty" xml:"mozi_request_id,omitempty"`
	// 如果success不为true，则自此段返回详细的错误信息
	ResponseMessage string `json:"response_message,omitempty" xml:"response_message,omitempty"`
	// 如果success为true，则返回0，否则此段返回详细的错误code
	ResponseCode string `json:"response_code,omitempty" xml:"response_code,omitempty"`
	// 是否操作成功,true代表操作成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclRoleRemoveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
	m.MoziRequestId = ""
	m.ResponseMessage = ""
	m.ResponseCode = ""
	m.IsSuccess = false
}

var poolAlibabaMoziAclRoleRemoveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclRoleRemoveAPIResponse)
	},
}

// GetAlibabaMoziAclRoleRemoveAPIResponse 从 sync.Pool 获取 AlibabaMoziAclRoleRemoveAPIResponse
func GetAlibabaMoziAclRoleRemoveAPIResponse() *AlibabaMoziAclRoleRemoveAPIResponse {
	return poolAlibabaMoziAclRoleRemoveAPIResponse.Get().(*AlibabaMoziAclRoleRemoveAPIResponse)
}

// ReleaseAlibabaMoziAclRoleRemoveAPIResponse 将 AlibabaMoziAclRoleRemoveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclRoleRemoveAPIResponse(v *AlibabaMoziAclRoleRemoveAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclRoleRemoveAPIResponse.Put(v)
}
