package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListuserrolesAPIResponse 查询并标记用户选择的角色 API返回值
// alibaba.campus.acl.new.listuserroles
//
// 查询并标记用户选择的角色
type AlibabaCampusAclNewListuserrolesAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewListuserrolesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewListuserrolesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewListuserrolesAPIResponseModel).Reset()
}

// AlibabaCampusAclNewListuserrolesAPIResponseModel is 查询并标记用户选择的角色 成功返回结果
type AlibabaCampusAclNewListuserrolesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_listuserroles_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewListuserrolesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewListuserrolesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewListuserrolesAPIResponse)
	},
}

// GetAlibabaCampusAclNewListuserrolesAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewListuserrolesAPIResponse
func GetAlibabaCampusAclNewListuserrolesAPIResponse() *AlibabaCampusAclNewListuserrolesAPIResponse {
	return poolAlibabaCampusAclNewListuserrolesAPIResponse.Get().(*AlibabaCampusAclNewListuserrolesAPIResponse)
}

// ReleaseAlibabaCampusAclNewListuserrolesAPIResponse 将 AlibabaCampusAclNewListuserrolesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewListuserrolesAPIResponse(v *AlibabaCampusAclNewListuserrolesAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewListuserrolesAPIResponse.Put(v)
}
