package moziacl

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziAclAppGetpermisspkgsAPIResponse 分页获取应用的权限套餐 API返回值
// alibaba.mozi.acl.app.getpermisspkgs
//
// 分页查询应用下的权限套餐列表
type AlibabaMoziAclAppGetpermisspkgsAPIResponse struct {
	model.CommonResponse
	AlibabaMoziAclAppGetpermisspkgsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaMoziAclAppGetpermisspkgsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaMoziAclAppGetpermisspkgsAPIResponseModel).Reset()
}

// AlibabaMoziAclAppGetpermisspkgsAPIResponseModel is 分页获取应用的权限套餐 成功返回结果
type AlibabaMoziAclAppGetpermisspkgsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mozi_acl_app_getpermisspkgs_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 获取应用的权限套餐列表结果对象
	Result *AppPermissionPackageResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaMoziAclAppGetpermisspkgsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaMoziAclAppGetpermisspkgsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaMoziAclAppGetpermisspkgsAPIResponse)
	},
}

// GetAlibabaMoziAclAppGetpermisspkgsAPIResponse 从 sync.Pool 获取 AlibabaMoziAclAppGetpermisspkgsAPIResponse
func GetAlibabaMoziAclAppGetpermisspkgsAPIResponse() *AlibabaMoziAclAppGetpermisspkgsAPIResponse {
	return poolAlibabaMoziAclAppGetpermisspkgsAPIResponse.Get().(*AlibabaMoziAclAppGetpermisspkgsAPIResponse)
}

// ReleaseAlibabaMoziAclAppGetpermisspkgsAPIResponse 将 AlibabaMoziAclAppGetpermisspkgsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaMoziAclAppGetpermisspkgsAPIResponse(v *AlibabaMoziAclAppGetpermisspkgsAPIResponse) {
	v.Reset()
	poolAlibabaMoziAclAppGetpermisspkgsAPIResponse.Put(v)
}
