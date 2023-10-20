package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclCancelpermiitemfromroleAPIResponse 取消角色和权限之间的关系 API返回值
// alibaba.campus.acl.cancelpermiitemfromrole
//
// 取消角色和权限之间的关系
type AlibabaCampusAclCancelpermiitemfromroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclCancelpermiitemfromroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclCancelpermiitemfromroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclCancelpermiitemfromroleAPIResponseModel).Reset()
}

// AlibabaCampusAclCancelpermiitemfromroleAPIResponseModel is 取消角色和权限之间的关系 成功返回结果
type AlibabaCampusAclCancelpermiitemfromroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_cancelpermiitemfromrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclCancelpermiitemfromroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclCancelpermiitemfromroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclCancelpermiitemfromroleAPIResponse)
	},
}

// GetAlibabaCampusAclCancelpermiitemfromroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclCancelpermiitemfromroleAPIResponse
func GetAlibabaCampusAclCancelpermiitemfromroleAPIResponse() *AlibabaCampusAclCancelpermiitemfromroleAPIResponse {
	return poolAlibabaCampusAclCancelpermiitemfromroleAPIResponse.Get().(*AlibabaCampusAclCancelpermiitemfromroleAPIResponse)
}

// ReleaseAlibabaCampusAclCancelpermiitemfromroleAPIResponse 将 AlibabaCampusAclCancelpermiitemfromroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclCancelpermiitemfromroleAPIResponse(v *AlibabaCampusAclCancelpermiitemfromroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclCancelpermiitemfromroleAPIResponse.Put(v)
}
