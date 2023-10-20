package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewListusermenuAPIResponse 查询用户有权限的菜单树 API返回值
// alibaba.campus.acl.new.listusermenu
//
// 查询用户有权限的菜单树
type AlibabaCampusAclNewListusermenuAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewListusermenuAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewListusermenuAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewListusermenuAPIResponseModel).Reset()
}

// AlibabaCampusAclNewListusermenuAPIResponseModel is 查询用户有权限的菜单树 成功返回结果
type AlibabaCampusAclNewListusermenuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_listusermenu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ListResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewListusermenuAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewListusermenuAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewListusermenuAPIResponse)
	},
}

// GetAlibabaCampusAclNewListusermenuAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewListusermenuAPIResponse
func GetAlibabaCampusAclNewListusermenuAPIResponse() *AlibabaCampusAclNewListusermenuAPIResponse {
	return poolAlibabaCampusAclNewListusermenuAPIResponse.Get().(*AlibabaCampusAclNewListusermenuAPIResponse)
}

// ReleaseAlibabaCampusAclNewListusermenuAPIResponse 将 AlibabaCampusAclNewListusermenuAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewListusermenuAPIResponse(v *AlibabaCampusAclNewListusermenuAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewListusermenuAPIResponse.Put(v)
}
