package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewFreezeroleAPIResponse 冻结角色 API返回值
// alibaba.campus.acl.new.freezerole
//
// 冻结角色
type AlibabaCampusAclNewFreezeroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewFreezeroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewFreezeroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewFreezeroleAPIResponseModel).Reset()
}

// AlibabaCampusAclNewFreezeroleAPIResponseModel is 冻结角色 成功返回结果
type AlibabaCampusAclNewFreezeroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_freezerole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewFreezeroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewFreezeroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewFreezeroleAPIResponse)
	},
}

// GetAlibabaCampusAclNewFreezeroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewFreezeroleAPIResponse
func GetAlibabaCampusAclNewFreezeroleAPIResponse() *AlibabaCampusAclNewFreezeroleAPIResponse {
	return poolAlibabaCampusAclNewFreezeroleAPIResponse.Get().(*AlibabaCampusAclNewFreezeroleAPIResponse)
}

// ReleaseAlibabaCampusAclNewFreezeroleAPIResponse 将 AlibabaCampusAclNewFreezeroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewFreezeroleAPIResponse(v *AlibabaCampusAclNewFreezeroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewFreezeroleAPIResponse.Put(v)
}
