package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewCheckuserpermissionAPIResponse 校验用户是否有权限 API返回值
// alibaba.campus.acl.new.checkuserpermission
//
// 校验用户是否有权限
type AlibabaCampusAclNewCheckuserpermissionAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewCheckuserpermissionAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewCheckuserpermissionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewCheckuserpermissionAPIResponseModel).Reset()
}

// AlibabaCampusAclNewCheckuserpermissionAPIResponseModel is 校验用户是否有权限 成功返回结果
type AlibabaCampusAclNewCheckuserpermissionAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_checkuserpermission_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewCheckuserpermissionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewCheckuserpermissionAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewCheckuserpermissionAPIResponse)
	},
}

// GetAlibabaCampusAclNewCheckuserpermissionAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewCheckuserpermissionAPIResponse
func GetAlibabaCampusAclNewCheckuserpermissionAPIResponse() *AlibabaCampusAclNewCheckuserpermissionAPIResponse {
	return poolAlibabaCampusAclNewCheckuserpermissionAPIResponse.Get().(*AlibabaCampusAclNewCheckuserpermissionAPIResponse)
}

// ReleaseAlibabaCampusAclNewCheckuserpermissionAPIResponse 将 AlibabaCampusAclNewCheckuserpermissionAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewCheckuserpermissionAPIResponse(v *AlibabaCampusAclNewCheckuserpermissionAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewCheckuserpermissionAPIResponse.Put(v)
}
