package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewSaverolewithmenuAPIResponse 保存角色级联保存角色和权限的关系 API返回值
// alibaba.campus.acl.new.saverolewithmenu
//
// 保存角色级联保存角色和权限的关系
type AlibabaCampusAclNewSaverolewithmenuAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewSaverolewithmenuAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewSaverolewithmenuAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewSaverolewithmenuAPIResponseModel).Reset()
}

// AlibabaCampusAclNewSaverolewithmenuAPIResponseModel is 保存角色级联保存角色和权限的关系 成功返回结果
type AlibabaCampusAclNewSaverolewithmenuAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_saverolewithmenu_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回结果
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewSaverolewithmenuAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewSaverolewithmenuAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewSaverolewithmenuAPIResponse)
	},
}

// GetAlibabaCampusAclNewSaverolewithmenuAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewSaverolewithmenuAPIResponse
func GetAlibabaCampusAclNewSaverolewithmenuAPIResponse() *AlibabaCampusAclNewSaverolewithmenuAPIResponse {
	return poolAlibabaCampusAclNewSaverolewithmenuAPIResponse.Get().(*AlibabaCampusAclNewSaverolewithmenuAPIResponse)
}

// ReleaseAlibabaCampusAclNewSaverolewithmenuAPIResponse 将 AlibabaCampusAclNewSaverolewithmenuAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewSaverolewithmenuAPIResponse(v *AlibabaCampusAclNewSaverolewithmenuAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewSaverolewithmenuAPIResponse.Put(v)
}
