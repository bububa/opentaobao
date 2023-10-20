package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclQueryallemppermiitemAPIResponse 查询员工全部权限(包括角色下面的权限) API返回值
// alibaba.campus.acl.queryallemppermiitem
//
// 查询员工全部权限(包括角色下面的权限)
type AlibabaCampusAclQueryallemppermiitemAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclQueryallemppermiitemAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclQueryallemppermiitemAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclQueryallemppermiitemAPIResponseModel).Reset()
}

// AlibabaCampusAclQueryallemppermiitemAPIResponseModel is 查询员工全部权限(包括角色下面的权限) 成功返回结果
type AlibabaCampusAclQueryallemppermiitemAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_queryallemppermiitem_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclQueryallemppermiitemAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclQueryallemppermiitemAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclQueryallemppermiitemAPIResponse)
	},
}

// GetAlibabaCampusAclQueryallemppermiitemAPIResponse 从 sync.Pool 获取 AlibabaCampusAclQueryallemppermiitemAPIResponse
func GetAlibabaCampusAclQueryallemppermiitemAPIResponse() *AlibabaCampusAclQueryallemppermiitemAPIResponse {
	return poolAlibabaCampusAclQueryallemppermiitemAPIResponse.Get().(*AlibabaCampusAclQueryallemppermiitemAPIResponse)
}

// ReleaseAlibabaCampusAclQueryallemppermiitemAPIResponse 将 AlibabaCampusAclQueryallemppermiitemAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclQueryallemppermiitemAPIResponse(v *AlibabaCampusAclQueryallemppermiitemAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclQueryallemppermiitemAPIResponse.Put(v)
}
