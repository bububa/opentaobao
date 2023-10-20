package campus

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCampusAclUpdategrantroletouserAPIResponse 修改用户到角色关系 API返回值
// alibaba.campus.acl.updategrantroletouser
//
// 修改用户到角色关系
type AlibabaCampusAclUpdategrantroletouserAPIResponse struct {
	model.CommonResponse
	AlibabaCampusAclUpdategrantroletouserAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaCampusAclUpdategrantroletouserAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaCampusAclUpdategrantroletouserAPIResponseModel).Reset()
}

// AlibabaCampusAclUpdategrantroletouserAPIResponseModel is 修改用户到角色关系 成功返回结果
type AlibabaCampusAclUpdategrantroletouserAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_campus_acl_updategrantroletouser_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *PojoResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaCampusAclUpdategrantroletouserAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaCampusAclUpdategrantroletouserAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaCampusAclUpdategrantroletouserAPIResponse)
	},
}

// GetAlibabaCampusAclUpdategrantroletouserAPIResponse 从 sync.Pool 获取 AlibabaCampusAclUpdategrantroletouserAPIResponse
func GetAlibabaCampusAclUpdategrantroletouserAPIResponse() *AlibabaCampusAclUpdategrantroletouserAPIResponse {
	return poolAlibabaCampusAclUpdategrantroletouserAPIResponse.Get().(*AlibabaCampusAclUpdategrantroletouserAPIResponse)
}

// ReleaseAlibabaCampusAclUpdategrantroletouserAPIResponse 将 AlibabaCampusAclUpdategrantroletouserAPIResponse 保存到 sync.Pool
func ReleaseAlibabaCampusAclUpdategrantroletouserAPIResponse(v *AlibabaCampusAclUpdategrantroletouserAPIResponse) {
	v.Reset()
	poolAlibabaCampusAclUpdategrantroletouserAPIResponse.Put(v)
}
