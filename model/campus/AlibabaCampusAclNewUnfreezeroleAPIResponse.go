package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclNewUnfreezeroleAPIResponse 解冻角色 API返回值
// alibaba.campus.acl.new.unfreezerole
//
// 解冻角色
type AlibabaCampusAclNewUnfreezeroleAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclNewUnfreezeroleAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewUnfreezeroleAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclNewUnfreezeroleAPIResponseModel).Reset()
}

// AlibabaCampusAclNewUnfreezeroleAPIResponseModel is 解冻角色 成功返回结果
type AlibabaCampusAclNewUnfreezeroleAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_new_unfreezerole_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclNewUnfreezeroleAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclNewUnfreezeroleAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclNewUnfreezeroleAPIResponse)
	},
}

// GetAlibabaCampusAclNewUnfreezeroleAPIResponse 从 sync.Pool 获取 AlibabaCampusAclNewUnfreezeroleAPIResponse
func GetAlibabaCampusAclNewUnfreezeroleAPIResponse() *AlibabaCampusAclNewUnfreezeroleAPIResponse {
	return poolAlibabaCampusAclNewUnfreezeroleAPIResponse.Get().(*AlibabaCampusAclNewUnfreezeroleAPIResponse)
}

// ReleaseAlibabaCampusAclNewUnfreezeroleAPIResponse 将 AlibabaCampusAclNewUnfreezeroleAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclNewUnfreezeroleAPIResponse(v *AlibabaCampusAclNewUnfreezeroleAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclNewUnfreezeroleAPIResponse.Put(v)
}
