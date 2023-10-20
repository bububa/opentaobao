package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewDeleteuserroleAPIResponse 删除管理员 API返回值
// alibaba.campus.acl.new.deleteuserrole
//
// 删除管理员
type AlibabaCampusAclNewDeleteuserroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewDeleteuserroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewDeleteuserroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewDeleteuserroleAPIResponseModel).Reset()
}

// AlibabaCampusAclNewDeleteuserroleAPIResponseModel is 删除管理员 成功返回结果
type AlibabaCampusAclNewDeleteuserroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_deleteuserrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewDeleteuserroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewDeleteuserroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewDeleteuserroleAPIResponse)
	},
}

// GetAlibabaCampusAclNewDeleteuserroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewDeleteuserroleAPIResponse
func GetAlibabaCampusAclNewDeleteuserroleAPIResponse() *AlibabaCampusAclNewDeleteuserroleAPIResponse {
	return poolAlibabaCampusAclNewDeleteuserroleAPIResponse.Get().(*AlibabaCampusAclNewDeleteuserroleAPIResponse)
}

// ReleaseAlibabaCampusAclNewDeleteuserroleAPIResponse 将 AlibabaCampusAclNewDeleteuserroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewDeleteuserroleAPIResponse(v *AlibabaCampusAclNewDeleteuserroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewDeleteuserroleAPIResponse.Put(v)
}
