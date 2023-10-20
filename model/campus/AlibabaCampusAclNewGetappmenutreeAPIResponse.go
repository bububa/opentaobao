package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewGetappmenutreeAPIResponse 查询应用下的菜单树 API返回值
// alibaba.campus.acl.new.getappmenutree
//
// 查询应用下的菜单树
type AlibabaCampusAclNewGetappmenutreeAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewGetappmenutreeAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewGetappmenutreeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewGetappmenutreeAPIResponseModel).Reset()
}

// AlibabaCampusAclNewGetappmenutreeAPIResponseModel is 查询应用下的菜单树 成功返回结果
type AlibabaCampusAclNewGetappmenutreeAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_getappmenutree_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewGetappmenutreeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewGetappmenutreeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewGetappmenutreeAPIResponse)
	},
}

// GetAlibabaCampusAclNewGetappmenutreeAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewGetappmenutreeAPIResponse
func GetAlibabaCampusAclNewGetappmenutreeAPIResponse() *AlibabaCampusAclNewGetappmenutreeAPIResponse {
	return poolAlibabaCampusAclNewGetappmenutreeAPIResponse.Get().(*AlibabaCampusAclNewGetappmenutreeAPIResponse)
}

// ReleaseAlibabaCampusAclNewGetappmenutreeAPIResponse 将 AlibabaCampusAclNewGetappmenutreeAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewGetappmenutreeAPIResponse(v *AlibabaCampusAclNewGetappmenutreeAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewGetappmenutreeAPIResponse.Put(v)
}
