package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewCheckusermenuAPIResponse 校验用户是否有菜单权限 API返回值
// alibaba.campus.acl.new.checkusermenu
//
// 校验用户是否有菜单权限
type AlibabaCampusAclNewCheckusermenuAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewCheckusermenuAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewCheckusermenuAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewCheckusermenuAPIResponseModel).Reset()
}

// AlibabaCampusAclNewCheckusermenuAPIResponseModel is 校验用户是否有菜单权限 成功返回结果
type AlibabaCampusAclNewCheckusermenuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_checkusermenu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewCheckusermenuAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewCheckusermenuAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewCheckusermenuAPIResponse)
	},
}

// GetAlibabaCampusAclNewCheckusermenuAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewCheckusermenuAPIResponse
func GetAlibabaCampusAclNewCheckusermenuAPIResponse() *AlibabaCampusAclNewCheckusermenuAPIResponse {
	return poolAlibabaCampusAclNewCheckusermenuAPIResponse.Get().(*AlibabaCampusAclNewCheckusermenuAPIResponse)
}

// ReleaseAlibabaCampusAclNewCheckusermenuAPIResponse 将 AlibabaCampusAclNewCheckusermenuAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewCheckusermenuAPIResponse(v *AlibabaCampusAclNewCheckusermenuAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewCheckusermenuAPIResponse.Put(v)
}
