package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListrolesAPIResponse 查询全部角色 API返回值
// alibaba.campus.acl.new.listroles
//
// 查询全部角色
type AlibabaCampusAclNewListrolesAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewListrolesAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewListrolesAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewListrolesAPIResponseModel).Reset()
}

// AlibabaCampusAclNewListrolesAPIResponseModel is 查询全部角色 成功返回结果
type AlibabaCampusAclNewListrolesAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_listroles_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewListrolesAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewListrolesAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewListrolesAPIResponse)
	},
}

// GetAlibabaCampusAclNewListrolesAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewListrolesAPIResponse
func GetAlibabaCampusAclNewListrolesAPIResponse() *AlibabaCampusAclNewListrolesAPIResponse {
	return poolAlibabaCampusAclNewListrolesAPIResponse.Get().(*AlibabaCampusAclNewListrolesAPIResponse)
}

// ReleaseAlibabaCampusAclNewListrolesAPIResponse 将 AlibabaCampusAclNewListrolesAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewListrolesAPIResponse(v *AlibabaCampusAclNewListrolesAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewListrolesAPIResponse.Put(v)
}
