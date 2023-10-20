package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclGrantpermiitemtoroleAPIResponse 权限赋予角色 API返回值
// alibaba.campus.acl.grantpermiitemtorole
//
// 权限赋予角色
type AlibabaCampusAclGrantpermiitemtoroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclGrantpermiitemtoroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel).Reset()
}

// AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel is 权限赋予角色 成功返回结果
type AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_grantpermiitemtorole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclGrantpermiitemtoroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclGrantpermiitemtoroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclGrantpermiitemtoroleAPIResponse)
	},
}

// GetAlibabaCampusAclGrantpermiitemtoroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclGrantpermiitemtoroleAPIResponse
func GetAlibabaCampusAclGrantpermiitemtoroleAPIResponse() *AlibabaCampusAclGrantpermiitemtoroleAPIResponse {
	return poolAlibabaCampusAclGrantpermiitemtoroleAPIResponse.Get().(*AlibabaCampusAclGrantpermiitemtoroleAPIResponse)
}

// ReleaseAlibabaCampusAclGrantpermiitemtoroleAPIResponse 将 AlibabaCampusAclGrantpermiitemtoroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclGrantpermiitemtoroleAPIResponse(v *AlibabaCampusAclGrantpermiitemtoroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclGrantpermiitemtoroleAPIResponse.Put(v)
}
