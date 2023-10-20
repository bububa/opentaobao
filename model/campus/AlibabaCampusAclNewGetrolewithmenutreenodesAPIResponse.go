package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse 根据角色id查询权限 API返回值
// alibaba.campus.acl.new.getrolewithmenutreenodes
//
// 根据角色id查询权限
type AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponseModel).Reset()
}

// AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponseModel is 根据角色id查询权限 成功返回结果
type AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_getrolewithmenutreenodes_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse)
	},
}

// GetAlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse
func GetAlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse() *AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse {
	return poolAlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse.Get().(*AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse)
}

// ReleaseAlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse 将 AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse(v *AlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewGetrolewithmenutreenodesAPIResponse.Put(v)
}
