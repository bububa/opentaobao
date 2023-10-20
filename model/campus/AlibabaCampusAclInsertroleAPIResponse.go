package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclInsertroleAPIResponse 新增角色 API返回值
// alibaba.campus.acl.insertrole
//
// 新增角色
type AlibabaCampusAclInsertroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclInsertroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclInsertroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclInsertroleAPIResponseModel).Reset()
}

// AlibabaCampusAclInsertroleAPIResponseModel is 新增角色 成功返回结果
type AlibabaCampusAclInsertroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_insertrole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *RoleRsp `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclInsertroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclInsertroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclInsertroleAPIResponse)
	},
}

// GetAlibabaCampusAclInsertroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclInsertroleAPIResponse
func GetAlibabaCampusAclInsertroleAPIResponse() *AlibabaCampusAclInsertroleAPIResponse {
	return poolAlibabaCampusAclInsertroleAPIResponse.Get().(*AlibabaCampusAclInsertroleAPIResponse)
}

// ReleaseAlibabaCampusAclInsertroleAPIResponse 将 AlibabaCampusAclInsertroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclInsertroleAPIResponse(v *AlibabaCampusAclInsertroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclInsertroleAPIResponse.Put(v)
}
