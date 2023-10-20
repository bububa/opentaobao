package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclQueryallroleAPIResponse 查询全部角色 API返回值
// alibaba.campus.acl.queryallrole
//
// 查询全部园区
type AlibabaCampusAclQueryallroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclQueryallroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclQueryallroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclQueryallroleAPIResponseModel).Reset()
}

// AlibabaCampusAclQueryallroleAPIResponseModel is 查询全部角色 成功返回结果
type AlibabaCampusAclQueryallroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_queryallrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CollectionResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclQueryallroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclQueryallroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclQueryallroleAPIResponse)
	},
}

// GetAlibabaCampusAclQueryallroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclQueryallroleAPIResponse
func GetAlibabaCampusAclQueryallroleAPIResponse() *AlibabaCampusAclQueryallroleAPIResponse {
	return poolAlibabaCampusAclQueryallroleAPIResponse.Get().(*AlibabaCampusAclQueryallroleAPIResponse)
}

// ReleaseAlibabaCampusAclQueryallroleAPIResponse 将 AlibabaCampusAclQueryallroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclQueryallroleAPIResponse(v *AlibabaCampusAclQueryallroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclQueryallroleAPIResponse.Put(v)
}
