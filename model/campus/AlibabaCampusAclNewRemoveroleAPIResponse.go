package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewRemoveroleAPIResponse 删除角色 API返回值
// alibaba.campus.acl.new.removerole
//
// 删除角色
type AlibabaCampusAclNewRemoveroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewRemoveroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewRemoveroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewRemoveroleAPIResponseModel).Reset()
}

// AlibabaCampusAclNewRemoveroleAPIResponseModel is 删除角色 成功返回结果
type AlibabaCampusAclNewRemoveroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_removerole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// {}
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewRemoveroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewRemoveroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewRemoveroleAPIResponse)
	},
}

// GetAlibabaCampusAclNewRemoveroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewRemoveroleAPIResponse
func GetAlibabaCampusAclNewRemoveroleAPIResponse() *AlibabaCampusAclNewRemoveroleAPIResponse {
	return poolAlibabaCampusAclNewRemoveroleAPIResponse.Get().(*AlibabaCampusAclNewRemoveroleAPIResponse)
}

// ReleaseAlibabaCampusAclNewRemoveroleAPIResponse 将 AlibabaCampusAclNewRemoveroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewRemoveroleAPIResponse(v *AlibabaCampusAclNewRemoveroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewRemoveroleAPIResponse.Put(v)
}
