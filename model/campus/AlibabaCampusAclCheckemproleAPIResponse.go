package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclCheckemproleAPIResponse 校验用户是否有该角色 API返回值
// alibaba.campus.acl.checkemprole
//
// 校验用户是否有该权限
type AlibabaCampusAclCheckemproleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclCheckemproleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclCheckemproleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclCheckemproleAPIResponseModel).Reset()
}

// AlibabaCampusAclCheckemproleAPIResponseModel is 校验用户是否有该角色 成功返回结果
type AlibabaCampusAclCheckemproleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_checkemprole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclCheckemproleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclCheckemproleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclCheckemproleAPIResponse)
	},
}

// GetAlibabaCampusAclCheckemproleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclCheckemproleAPIResponse
func GetAlibabaCampusAclCheckemproleAPIResponse() *AlibabaCampusAclCheckemproleAPIResponse {
	return poolAlibabaCampusAclCheckemproleAPIResponse.Get().(*AlibabaCampusAclCheckemproleAPIResponse)
}

// ReleaseAlibabaCampusAclCheckemproleAPIResponse 将 AlibabaCampusAclCheckemproleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclCheckemproleAPIResponse(v *AlibabaCampusAclCheckemproleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclCheckemproleAPIResponse.Put(v)
}
