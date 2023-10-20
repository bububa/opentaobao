package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclCancelrolesfromuserAPIResponse 撤销用户授予的角色 API返回值
// alibaba.campus.acl.cancelrolesfromuser
//
// 撤销用户授予的角色
type AlibabaCampusAclCancelrolesfromuserAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclCancelrolesfromuserAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclCancelrolesfromuserAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclCancelrolesfromuserAPIResponseModel).Reset()
}

// AlibabaCampusAclCancelrolesfromuserAPIResponseModel is 撤销用户授予的角色 成功返回结果
type AlibabaCampusAclCancelrolesfromuserAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_cancelrolesfromuser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclCancelrolesfromuserAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclCancelrolesfromuserAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclCancelrolesfromuserAPIResponse)
	},
}

// GetAlibabaCampusAclCancelrolesfromuserAPIResponse 从 sync.Pool 获取 AlibabaCampusAclCancelrolesfromuserAPIResponse
func GetAlibabaCampusAclCancelrolesfromuserAPIResponse() *AlibabaCampusAclCancelrolesfromuserAPIResponse {
	return poolAlibabaCampusAclCancelrolesfromuserAPIResponse.Get().(*AlibabaCampusAclCancelrolesfromuserAPIResponse)
}

// ReleaseAlibabaCampusAclCancelrolesfromuserAPIResponse 将 AlibabaCampusAclCancelrolesfromuserAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclCancelrolesfromuserAPIResponse(v *AlibabaCampusAclCancelrolesfromuserAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclCancelrolesfromuserAPIResponse.Put(v)
}
